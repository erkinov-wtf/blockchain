package chain

import (
	"blockchain/block"
	"blockchain/block/types"
)

type Blockchain struct {
	Blocks []*types.Block
}

func (bc *Blockchain) AddBlock(data string) {
	previousBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := block.NewBlock(data, previousBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*types.Block{block.NewGenesisBlock()}}
}
