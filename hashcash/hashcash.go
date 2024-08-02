package hashcash

import (
	"blockchain/block/types"
	"blockchain/utils"
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

const targetBits = 15

type ProofOfWork struct {
	Block  *types.Block
	Target *big.Int
}

func NewProofOfWork(block *types.Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))
	pow := &ProofOfWork{block, target}

	return pow
}

func (pow *ProofOfWork) PrepareData(counter int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevBlockHash,
			pow.Block.Data,
			utils.IntToHex(pow.Block.Timestamp),
			utils.IntToHex(int64(targetBits)),
			utils.IntToHex(int64(counter)),
		},
		[]byte{},
	)

	return data
}

func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	maxCounter := math.MaxInt64
	counter := 0

	fmt.Printf("Mining the block containing \"%s\"\n", pow.Block.Data)
	for counter < maxCounter {
		data := pow.PrepareData(counter)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.Target) == -1 {
			break
		} else {
			counter++
		}
	}

	fmt.Print("\n\n")
	return counter, hash[:]
}

func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.PrepareData(pow.Block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.Target) == -1

	return isValid
}
