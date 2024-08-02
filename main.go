package main

import (
	"blockchain/chain"
	"fmt"
)

func main() {
	bc := chain.NewBlockchain()

	bc.AddBlock("Second Block: Sent some coins")
	bc.AddBlock("Third Block: Sent other coins")

	for _, block := range bc.Blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Previous Hash: %x\n", block.Hash)
		fmt.Printf("Previous Hash: %s\n", block.Data)
	}
}
