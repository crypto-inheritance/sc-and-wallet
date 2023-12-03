## Setting Up
1. Solidity contract to ABI
```
solc --abi contracts/Inheritance.sol -o contracts
```

2. ABI to Go package
```
abigen --abi contracts/Inheritance.abi --pkg contract_pkg --out contract_pkg/inheritance.go
```

3. Dependency installation

```
go get github.com/ethereum/go-ethereum
go get github.com/ethereum/go-ethereum/accounts/abi/bind

go get gopkg.in/mgo.v2
go get github.com/BurntSushi/toml
```


