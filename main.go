package main

import (
	"bitcoin-chain/blockChain"
	"fmt"
)

func main() {
	bc := blockChain.NewBlockChain()
	bc.AddBlock("hello1,hello2")
	bc.AddBlock("hello3,hello4")

	for _, block := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
	fmt.Println("bitcoin chain start running !!")
}
