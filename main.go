package main

import (
	"bitcoin-chain/blockChain"
	"bitcoin-chain/cli"
)

func main() {

	bc := blockChain.NewBlockchain()
	defer bc.Db.Close()

	cli := cli.CLI{bc}
	cli.Run()
}
