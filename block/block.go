package block

import (
	"blockchain/block/types"
	"blockchain/hashcash"
	"time"
)

func NewGenesisBlock() *types.Block {
	return NewBlock("Genesis Block", []byte{})
}

func NewBlock(data string, prevBlockHash []byte) *types.Block {
	block := &types.Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := hashcash.NewProofOfWork(block)
	counter, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = counter

	return block
}
