package blockChain

import (
	"time"
)

type Block struct {
	Timestamp     int64  // 毫秒时间戳
	Data          []byte // 区块数据
	PrevBlockHash []byte // 上一个区块hash
	Hash          []byte // 当前区块hash
	Nonce         int
}

// 创建区块
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().UnixMilli(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	block.Nonce = nonce
	block.Hash = hash[:]
	return block
}

//创建创世区块
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
