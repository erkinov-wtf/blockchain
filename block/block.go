package block

import (
	"blockchain/block/types"
	"blockchain/hashcash"
	"blockchain/transaction"
	"time"
)

func NewGenesisBlock(coinbase *transaction.Transaction) *types.Block {
	return NewBlock([]*transaction.Transaction{coinbase}, []byte{})
}

func NewBlock(transactions []*transaction.Transaction, prevBlockHash []byte) *types.Block {
	block := &types.Block{time.Now().Unix(), transactions, prevBlockHash, []byte{}, 0}
	pow := hashcash.NewProofOfWork(block)
	counter, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = counter

	return block
}
