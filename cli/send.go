package cli

import (
	"blockchain/chain"
	"blockchain/transaction"
	"blockchain/wallet"
	"fmt"
	"log"
)

func (cli *CLI) send(from, to string, amount int) {
	if !wallet.ValidateAddress(from) {
		log.Panic("ERROR: Sender address is not valid")
	}
	if !wallet.ValidateAddress(to) {
		log.Panic("ERROR: Recipient address is not valid")
	}

	bc := chain.NewBlockchain(from)
	defer bc.DB.Close()

	tx := chain.NewUTXOTransaction(from, to, amount, bc)
	bc.MineBlock([]*transaction.Transaction{tx})
	fmt.Println("Success!")
}
