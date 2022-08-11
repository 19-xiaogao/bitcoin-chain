package blockChain

type BlockChain struct {
	Blocks []*Block
}

func (bc *BlockChain) AddBlock(data string) {
	prvBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data, prvBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}
