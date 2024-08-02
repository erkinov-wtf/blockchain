package main

import (
	"blockchain/chain"
	"blockchain/cli"
)

func main() {
	bc := chain.NewBlockchain()
	defer bc.DB.Close()

	cli := cli.CLI{BC: bc}
	cli.Run()
}
