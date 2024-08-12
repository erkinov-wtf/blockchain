package cli

import (
	"blockchain/chain"
	"blockchain/wallet"
	"fmt"
	"log"
)

func (cli *CLI) createBlockchain(address string) {
	if !wallet.ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	bc := chain.CreateBlockchain(address)
	bc.DB.Close()
	fmt.Println("Done!")
}
