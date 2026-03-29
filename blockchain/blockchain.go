package blockchain

import "sync"

type Blockchain struct {
	mu         sync.RWMutex
	Blocks     []Block
	Difficulty int
}

func NewBlockchain() *Blockchain {
	difficulty := 3
	genesis := NewBlock(0, []Transaction{}, "", difficulty)

	return &Blockchain{
		Blocks:     []Block{genesis},
		Difficulty: difficulty,
	}
}

func (bc *Blockchain) GetLatestBlock() Block {
	bc.mu.RLock()
	defer bc.mu.RUnlock()

	return bc.Blocks[len(bc.Blocks)-1]
}

func (bc *Blockchain) AddBlock(txs []Transaction) Block {
	bc.mu.Lock()
	defer bc.mu.Unlock()

	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(prevBlock.Index+1, txs, prevBlock.Hash, bc.Difficulty)

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
