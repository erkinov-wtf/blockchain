package cli

import (
	"blockchain/chain"
	"blockchain/utils"
	"blockchain/wallet"
	"fmt"
	"log"
)

func (cli *CLI) getBalance(address string) {
	if !wallet.ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	bc := chain.NewBlockchain(address)
	defer bc.DB.Close()

	balance := 0
	pubKeyHash := utils.Base58Decode([]byte(address))
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
	UTXOs := bc.FindUTXO(pubKeyHash)

	for _, out := range UTXOs {
		balance += out.Value
	}

	fmt.Printf("Balance of '%s': %d\n", address, balance)
}
