package main

import (
	"context"
	"fmt"

	inh "main/contract_pkg"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	// "github.com/miguelmota/go-web3-example/greeter"
	"log"
)

var eventSignatureList = []common.Hash{crypto.Keccak256Hash([]byte("SignatoryAdded(address)")),
	crypto.Keccak256Hash([]byte("SignatoryApproved(address)")),
	crypto.Keccak256Hash([]byte("SignatoryRemoved(address)")),
	crypto.Keccak256Hash([]byte("WeightsUpdated(address,uint)")),
	crypto.Keccak256Hash([]byte("TransactionRequested(uint,address,uint)")),
	crypto.Keccak256Hash([]byte("ApprovalReceived(uint,address)")),
	crypto.Keccak256Hash([]byte("TransactionApproved(uint)")),
	crypto.Keccak256Hash([]byte("AssetsDeposited(address,uint)")),
	crypto.Keccak256Hash([]byte("InheritanceRequested(uint,address,uint)")),
	crypto.Keccak256Hash([]byte("InheritanceRequestCanceled(uint)")),
	crypto.Keccak256Hash([]byte("InheritanceRequestApproved(uint)"))}

func getEventAction(contract *inh.Inheritance, topics []common.Hash) {
	switch topics[0] {
	case eventSignatureList[0]:
		fmt.Printf("State Changed: new signatory added... \n")
		break
	case eventSignatureList[1]:
		fmt.Printf("State Changed: new signatory approved... \n")
		break
	case eventSignatureList[2]:
		fmt.Printf("State Changed: signatory removed... \n")
		break
	case eventSignatureList[3]:
		fmt.Printf("State Changed: weights updated... \n")
		break
	case eventSignatureList[4]:
		fmt.Printf("State Changed: new transaction requested... \n")
		break
	case eventSignatureList[5]:
		fmt.Printf("State Changed: new transaction approval received.. \n")
		break
	case eventSignatureList[6]:
		fmt.Printf("State Changed: transaction approved... \n")
		break
	case eventSignatureList[7]:
		fmt.Printf("State Changed: assets deposited... \n")
		break
	case eventSignatureList[8]:
		fmt.Printf("State Changed: inheritance requested... \n")
		break
	case eventSignatureList[9]:
		fmt.Printf("State Changed: inheritance request canceled... \n")
		break
	case eventSignatureList[10]:
		fmt.Printf("State Changed: inheritance approved... \n")
		break
	default:
		fmt.Printf("State Changed...")
		break
	}

	getAndPrintInfo(contract)
}

func getAndPrintInfo(contract *inh.Inheritance) {
	signatories := make(map[common.Address]bool)

	balance, err := contract.Balance(nil)
	if err != nil {
		log.Fatal(err)
	}

	owner, err := contract.Owner(nil)
	if err != nil {
		log.Fatal(err)
	}
	signatories[owner] = true

	signatoryAddresses, err := contract.GetSignatoryAddresses(nil)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(signatoryAddresses); i++ {
		approved, err := contract.Signatories(nil, signatoryAddresses[i])
		if err != nil {
			log.Fatal(err)
		}
		signatories[signatoryAddresses[i]] = approved
	}

	approvalThreshold, err := contract.ApprovalThreshold(nil)
	if err != nil {
		log.Fatal(err)
	}

	approvedSignatoryCount, err := contract.ApprovedSignatoryCount(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("\n\n****************************************************************************************\n")
	fmt.Print("***                       Current configuration of the Wallet                        ***\n")
	fmt.Print("***                                                                                  ***\n")
	fmt.Print("***                                                                                  ***\n")
	fmt.Printf("***  Balance: %15d                                                        ***\n", balance)
	fmt.Printf("***  Transaction Approval Threshold (N/M): %d / %d                                     ***\n", approvalThreshold, approvedSignatoryCount)
	fmt.Print("***                                                                                  ***\n")
	for address, approved := range signatories {
		fmt.Printf("***  Address: %s;  Approved: %5t           ***\n", address.Hex(), approved)
	}
	fmt.Print("***                                                                                  ***\n")
	fmt.Print("****************************************************************************************\n")

}

func main() {
	// client, err := ethclient.Dial("https://sepolia.infura.io/v3/7f2e2b20cb9348adb5fa1a4af7be07b6")
	client, err := ethclient.Dial("wss://sepolia.infura.io/ws/v3/7f2e2b20cb9348adb5fa1a4af7be07b6")

	if err != nil {
		log.Fatal(err)
	}

	contractAddressStr := "0x542d67E8d5eCCF1919639366EF9aF312403Bc493"

	contractAddress := common.HexToAddress(contractAddressStr)
	inhClient, err := inh.NewInheritance(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	/*
		//Calling a function
		priv := "7ea39ba38cb7acc520f9a91564c50ec497b00ca9145dc83c59284cb44c24a617"
		key, err := crypto.HexToECDSA(priv)
		print("Key : %x\n", key)

		auth := bind.NewKeyedTransactor(key)

		tx, err := inhClient.AddSignatory(auth, "0x64eD3A517310763fd6EfD8D2CC0921baFbAbD286")

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Pending TX: 0x%x\n", tx.Hash())
	*/

	getAndPrintInfo(inhClient)

	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	var ch = make(chan types.Log)
	ctx := context.Background()

	sub, err := client.SubscribeFilterLogs(ctx, query, ch)
	if err != nil {
		log.Println("Subscribe:", err)
		return
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case log := <-ch:
			fmt.Printf(".\n.\n.\n")
			fmt.Println("Log: ", log)
			getEventAction(inhClient, log.Topics)
		}
	}

}
