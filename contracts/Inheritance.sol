/**
 *Submitted for verification at Etherscan.io on 2023-11-27
*/

// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

contract Inheritance {
    address public owner;
    mapping(address => bool) public signatories;
    int public approvedSignatoryCount = 0;
    address[] public signatoryAddresses;
    mapping(address => uint) public signatoryIndexes;
    mapping(address => uint) public signatoryWeights;
    uint private totalWeight = 0;
    uint public balance;

    event SignatoryAdded(address indexed signatory);
    event SignatoryApproved(address indexed signatory);
    event SignatoryRemoved(address indexed signatory);
    event WeightsUpdated(address indexed signatory, uint maxShares);

    constructor() {
        owner = msg.sender;
        signatories[owner] = true;
        approvedSignatoryCount += 1;
    }

    modifier onlyOwner() {
        require(msg.sender == owner, "Only the owner can perform this action.");
        _;
    }

    modifier onlySignatories() {
        require(signatories[msg.sender], "Only an approved signatory can perform this action.");
        _;
    }

    // Functions to manage the signatories
    modifier onlyPendingSignatories() {
        bool isPending = false;
        for (uint i = 0; i < signatoryAddresses.length; i++) {
            if (signatoryAddresses[i] == msg.sender && !signatories[msg.sender]) {
                isPending = true;
                break;
            }
        }
        require(isPending, "Only a pending signatory can perform this action.");
        _;
    }


    function addSignatory(address _signatory, uint _weight) public onlyOwner {
        require(signatoryIndexes[_signatory] == 0, "Signatory already added.");
        // require((totalWeight + _weight) <= 1000, "The total weight cannot be greater than 1000.");
        signatoryWeights[_signatory] = _weight;
        totalWeight += _weight;
        signatoryAddresses.push(_signatory);
        signatoryIndexes[_signatory] = signatoryAddresses.length;
        
        emit SignatoryAdded(_signatory);
    }

    function approveSignatory(address _signatory, uint _weight) public onlyOwner {
        require(!signatories[_signatory], "Signatory already approved.");
        if (signatoryIndexes[_signatory] == 0){
            signatoryAddresses.push(_signatory);
        }
        totalWeight = totalWeight - signatoryWeights[_signatory] + _weight;
        signatoryWeights[_signatory] = _weight;
        signatories[_signatory] = true;
        approvedSignatoryCount += 1;
        emit SignatoryApproved(_signatory);
    }

    function removeSignatory(address _signatory) public onlyOwner {
        require(signatoryIndexes[_signatory] != 0, "Signatory does not exist.");
        for(uint i = signatoryIndexes[_signatory]-1; i < signatoryAddresses.length - 1; i++){
            signatoryAddresses[i] = signatoryAddresses[i+1];
            signatoryIndexes[signatoryAddresses[i]] -= 1;
        }
        signatoryIndexes[_signatory] = 0;
        signatoryAddresses.pop();
        if (signatories[_signatory]){
            signatories[_signatory] = false;
            approvedSignatoryCount -= 1;
        }
        emit SignatoryRemoved(_signatory);
    }
    

    function getSignatoryAddresses() public view returns (address[] memory) {
        return signatoryAddresses;
    }

    function setWeightForSignatory(address _signatory, uint _weight) public onlyOwner {
        require(signatoryIndexes[_signatory] != 0, "Signatory not found.");
        totalWeight = totalWeight - signatoryWeights[_signatory] + _weight;
        signatoryWeights[_signatory] = _weight;
        emit WeightsUpdated(_signatory, _weight);
    }


    // Transaction management part
    struct Transaction {
        address signatory;
        uint amount;
        bool approved;
        uint approvalCount;
        mapping(address => bool) approvals;
    }

    Transaction[] public transactions;
    uint public approvalThreshold = 1;
    uint private lastTransactionTime = block.timestamp;
    address payable private walletAddress;
    bool public isWalletAddressSet = false;

    event TransactionRequested(uint indexed transactionId, address indexed signatory, uint amount);
    event ApprovalReceived(uint indexed transactionId, address indexed approver);
    event TransactionApproved(uint indexed transactionId);
    event AssetsDeposited(address indexed recipient, uint indexed amount);

    function getTransactionCount() public view returns (uint) {
        return transactions.length;
    }

    function getApprovedTransactionCount() public view returns (uint) {
        uint count = 0;
        for (uint i = 0; i < transactions.length; i++){
            if(transactions[i].approved){
                count++;
            }
        }
        return count;
    }
    function getLastPendingTransactionId() public view returns (uint){
        for (uint i = transactions.length - 1; i >= 0; i--){
            if(!transactions[i].approved){
                return i;
            }
        }
        return transactions.length;
    }

    function setApprovalThreshold(uint _threshold) public onlyOwner {
        approvalThreshold = _threshold;
    }

    function setLastTransactionTime(uint _time) public onlyOwner{
        lastTransactionTime = _time;
    }

    function setWalletAddress(address payable _walletAddress) public onlyOwner{
        walletAddress = _walletAddress;
        isWalletAddressSet = true;
    }

    function addBalance(uint amount) public onlyOwner{
        balance += amount;
        emit AssetsDeposited(walletAddress, amount);
    }


    function getWithdrawLimitForAddress(address _signatory) public view onlySignatories returns (uint) {
        if (_signatory == owner){
            return balance;
        }
        return (signatoryWeights[_signatory] * 1000 / totalWeight) * balance / 1000;
    }

    function getWeightForAmount(address _signatory, uint _amount) public view onlySignatories returns (uint) {
        return _amount * signatoryWeights[_signatory] / getWithdrawLimitForAddress(_signatory);
    }

    function depositFunds() payable public {
        require(msg.value > 0, "Must send some Ether.");
        require(isWalletAddressSet, "Wallet address is not set.");
        walletAddress.transfer(msg.value);
        balance += msg.value;
        lastTransactionTime = block.timestamp;
        emit AssetsDeposited(walletAddress, msg.value);
    }

    function requestTransaction(uint _amount) public onlySignatories {
        require(_amount <= getWithdrawLimitForAddress(msg.sender), "Amount exceeds allowed shares.");

        Transaction storage newTransaction = transactions.push();
        newTransaction.signatory = msg.sender;
        newTransaction.amount = _amount;
        newTransaction.approvalCount = 1;
        newTransaction.approved = false;
        newTransaction.approvals[msg.sender] = true;

        if (approvalThreshold == 1){
            newTransaction.approved = true;
            balance -= _amount;
            lastTransactionTime = block.timestamp;
            emit TransactionApproved(transactions.length - 1);
        }

        if (msg.sender != owner){
            signatoryWeights[msg.sender] -= getWeightForAmount(msg.sender, _amount);
        }
        emit TransactionRequested(transactions.length - 1, msg.sender, _amount);
    }

    function approveTransaction(uint _transactionId) public onlySignatories {
        require(_transactionId < transactions.length, "Invalid transaction ID.");
        Transaction storage transaction = transactions[_transactionId];
        require(!transaction.approvals[msg.sender], "Signatory has already approved this transaction.");

        transaction.approvals[msg.sender] = true;
        transaction.approvalCount++;
        emit ApprovalReceived(_transactionId, msg.sender);

        if (transaction.approvalCount >= approvalThreshold && !transaction.approved) {
            transaction.approved = true;
            balance -= transaction.amount;
            lastTransactionTime = block.timestamp;
            emit TransactionApproved(_transactionId);
        }
    }


    // Inheritance management 
    uint public inactivityThreshold = 90 days;
    uint public inheritanceDeadline = 90 days;   
    uint public inheritanceRequestTime = 0;    
    bool public inheritanceRequestActive = false;
    bool public inheritanceDeadlineActive = false;
    address public governmentAddress;

    modifier onlyGovernment() {
        require(msg.sender == governmentAddress, "Only verified government can perform this operation.");
        _;
    }

    
    event InheritanceRequested(uint indexed inheritanceRequestTime, address indexed signatory, uint deadline);
    event InheritanceRequestCanceled(uint indexed inheritanceRequestTime);
    event InheritanceRequestApproved(uint indexed inheritanceRequestTime);

    function setInactivityThreshold(uint _threshold) public onlyOwner {
        inactivityThreshold = _threshold;
    }
    
    function setInheritanceDeadline(uint _deadline) public onlyOwner {
        inheritanceDeadline = _deadline;
    }

    function setGovernmentAddress(address _address) public onlyOwner{
        governmentAddress = _address;
    }


    function requestInheritance() public onlyPendingSignatories {
        require(block.timestamp >= (lastTransactionTime + inactivityThreshold), "Inheritance request not valid yet.");
        require((!inheritanceRequestActive && !inheritanceDeadlineActive) || 
                (inheritanceRequestActive && !inheritanceDeadlineActive), "Inheritance is already requested.");
        inheritanceRequestActive = true;
        inheritanceDeadlineActive = true;
        inheritanceRequestTime = block.timestamp;
        emit InheritanceRequested(inheritanceRequestTime, msg.sender, inheritanceDeadline);
    }

    function cancelInheritanceRequest() public {
        require((msg.sender == owner) || (msg.sender == governmentAddress), "Only Owner and Government is allowed");
        require(inheritanceRequestActive && inheritanceDeadlineActive, "Inheritance request is not active.");
        inheritanceDeadlineActive = false;
        emit InheritanceRequestCanceled(inheritanceRequestTime);
    }

    function confirmInheritanceRequest() public onlyGovernment{
        require(inheritanceRequestActive && inheritanceDeadlineActive, "Inheritance request is not active.");
        for (uint i = 0; i < signatoryAddresses.length; i++) {
            signatories[signatoryAddresses[i]] = true;
            approvedSignatoryCount += 1;
        }
        approvalThreshold = 1;
        inheritanceRequestActive = false;
        emit InheritanceRequestApproved(inheritanceRequestTime);
    }

    function checkInheritanceRequestStatus() public onlyPendingSignatories returns (bool){
        require(inheritanceRequestActive, "Inheritance request is not active.");
        if(block.timestamp >= (inheritanceRequestTime + inheritanceDeadline)){
            for (uint i = 0; i < signatoryAddresses.length; i++) {
                signatories[signatoryAddresses[i]] = true;
                approvedSignatoryCount += 1;
            }
            approvalThreshold = 1;
            inheritanceRequestActive = false;
            emit InheritanceRequestApproved(inheritanceRequestTime);
            return true;
        }
        return false;
    }

}