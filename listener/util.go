package listener

import (
	"math/big"
)

// Config has the configuration data needed provided by toml file
type Config struct {
	ProviderURL     string
	ContractAddress string
	PrivateKey      string
	AbiPath         string
	EventName       string
	BlockNumber     *big.Int
	DbName          string
	DbCollection    string
	DbHost          string
}

// conf holds the filled Config struct
var Conf *Config
