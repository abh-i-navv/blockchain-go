package blockchain

import "sync"

type Blockchain struct {
	mu     sync.RWMutex
	Blocks []Block
}

func NewBlockchain() *Blockchain {
	genesis := NewBlock(0, "Genesis Block", "")

	return &Blockchain{
		Blocks: []Block{genesis},
	}
}

func (bc *Blockchain) GetLatestBlock() Block {
	bc.mu.RLock()
	defer bc.mu.RUnlock()

	return bc.Blocks[len(bc.Blocks)-1]
}

func (bc *Blockchain) AddBlock(data string) Block {
	bc.mu.Lock()
	defer bc.mu.Unlock()

	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(prevBlock.Index+1, data, prevBlock.Hash)

	bc.Blocks = append(bc.Blocks, newBlock)
	return newBlock
}

func (bc *Blockchain) IsValid() bool {
	bc.mu.RLock()
	defer bc.mu.RUnlock()

	for i := 1; i < len(bc.Blocks); i++ {
		curr := bc.Blocks[i]
		prev := bc.Blocks[i-1]

		if curr.Hash != curr.CalculateHash() {
			return false
		}

		if curr.PrevHash != prev.Hash {
			return false
		}
	}
	return true
}
