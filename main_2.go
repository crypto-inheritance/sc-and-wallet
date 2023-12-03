package main

import (
	"log"
	l "main/listener"

	"github.com/BurntSushi/toml"
)

// main function of the program
func main2() {
	if _, err := toml.DecodeFile("./config.toml", &l.Conf); err != nil {
		log.Fatal("Erro ", err)
	}
	provider := new(l.Provider)
	provider.SetUp()
	provider.ListenToEvent()
	//Just for testnet
	// var nonce int64 = 6
	// p.auth.Nonce = big.NewInt(nonce)
}
