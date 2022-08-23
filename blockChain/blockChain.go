package blockChain

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

type BlockChain struct {
	Tip []byte
	Db  *bolt.DB
}

const dbFile = "blockchain.db"
const blocksBucket = "blocks"

type Iterator struct {
	CurrentHash []byte
	Db          *bolt.DB
}

func (bc *BlockChain) AddBlock(data string) {
	var lastHash []byte
	err := bc.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lastHash = b.Get([]byte("l"))
		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	newBlock := NewBlock(data, lastHash)

	bc.Db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		err := b.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			log.Panic(err)
		}
		err = b.Put([]byte("l"), newBlock.Hash)
		if err != nil {
			log.Panic(err)
		}
		bc.Tip = newBlock.Hash
		return nil
	})
}

func (bc *BlockChain) Iterator() *Iterator {
	bci := &Iterator{CurrentHash: bc.Tip, Db: bc.Db}
	return bci
}

func (i *Iterator) Next() *Block {
	var block *Block
	err := i.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodeBlock := b.Get(i.CurrentHash)
		block = DeserializeBlock(encodeBlock)
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	i.CurrentHash = block.PrevBlockHash
	return block
}

func NewBlockchain() *BlockChain {
	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		if b == nil {
			fmt.Println("No exiting blockchain found, creating a new one ...")
			genesis := NewGenesisBlock()
			b, err := tx.CreateBucket([]byte(blocksBucket))
			if err != nil {
				log.Panic(err)
			}
			err = b.Put(genesis.Hash, genesis.Serialize())
			if err != nil {
				log.Panic(err)
			}
			err = b.Put([]byte("l"), genesis.Hash)
			if err != nil {
				log.Panic(err)
			}
			tip = genesis.Hash
		} else {
			tip = b.Get([]byte("l"))
		}
		return nil
	})

	if err != nil {
		log.Panic(err)
	}
	bc := BlockChain{tip, db}
	return &bc
}
