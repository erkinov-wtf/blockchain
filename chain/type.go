package chain

import "github.com/boltdb/bolt"

const dbFile = "storage/database/blockchain.db"
const blocksBucket = "blocks"

type Blockchain struct {
	Tip []byte
	DB  *bolt.DB
}

type BlockchainIterator struct {
	currentHash []byte
	DB          *bolt.DB
}
