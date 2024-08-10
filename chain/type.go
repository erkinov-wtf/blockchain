package chain

import "github.com/boltdb/bolt"

const dbFile = "storage/database/blockchain.db"
const blocksBucket = "blocks"
const genesisCoinbaseData = "In the ninja world those who break the rules and laws are regarded as scum, but those who would abandon even one of their friends are even worse than scum. - Obito Uchiha"

type Blockchain struct {
	Tip []byte
	DB  *bolt.DB
}

type BlockchainIterator struct {
	currentHash []byte
	DB          *bolt.DB
}
