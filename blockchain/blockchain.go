package blockchain

import "sync"

type Blockchain struct {
	mu         sync.RWMutex
	Blocks     []Block
	Difficulty int
	Storage    *Storage
}

func NewBlockchain() *Blockchain {
	difficulty := 2

	storage, err := NewStorage("./data")

	if err != nil {
		panic(err)
	}

	blocks, _ := storage.LoadBlocks()

	if len(blocks) == 0 {
		genesis := NewBlock(0, []Transaction{}, "", difficulty)

		return &Blockchain{
			Blocks:     []Block{genesis},
			Difficulty: difficulty,
			Storage:    storage,
		}
	}

	return &Blockchain{
		Blocks:     blocks,
		Difficulty: difficulty,
		Storage:    storage,
	}
}

func (bc *Blockchain) GetLatestBlock() Block {
	bc.mu.RLock()
	defer bc.mu.RUnlock()

	return bc.Blocks[len(bc.Blocks)-1]
}

func (bc *Blockchain) AddBlock(txs []Transaction, miner []byte) Block {
	bc.mu.Lock()
	defer bc.mu.Unlock()

	// reward to miner
	rewardTx := Transaction{
		From:   []byte("SYSTEM"),
		To:     miner,
		Amount: 50,
	}

	validTxs := []Transaction{rewardTx}

	tempBal := make(map[string]float64)
	tempBal[string(miner)] += 50

	for _, block := range bc.Blocks {
		for _, tx := range block.Transactions {
			from := string(tx.From)
			to := string(tx.To)

			if tempBal[from] < tx.Amount {
				continue
			}

			tempBal[from] -= tx.Amount
			tempBal[to] += tx.Amount

			if bc.IsValidTransactionUnsafe(tx) {
				validTxs = append(validTxs, tx)
			}
		}
	}

	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(prevBlock.Index+1, validTxs, prevBlock.Hash, bc.Difficulty)

	bc.Blocks = append(bc.Blocks, newBlock)
	bc.Storage.SaveBlock(newBlock)
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

func (bc *Blockchain) GetBalance(address []byte) float64 {
	bc.mu.RLock()
	defer bc.mu.RUnlock()

	return bc.GetBalanceUnsafe(address)
}

func (bc *Blockchain) GetBalanceUnsafe(address []byte) float64 {
	balance := 0.0

	for _, block := range bc.Blocks {
		for _, tx := range block.Transactions {
			if string(tx.From) == string(address) {
				balance -= tx.Amount
			}

			if string(tx.To) == string(address) {
				balance += tx.Amount
			}
		}
	}
	return balance
}

func (bc *Blockchain) IsValidTransaction(tx Transaction) bool {
	if !tx.Verify() {
		return false
	}

	balance := bc.GetBalance(tx.From)

	if balance < tx.Amount {
		return false
	}

	return true
}

func (bc *Blockchain) IsValidTransactionUnsafe(tx Transaction) bool {
	if !tx.Verify() {
		return false
	}

	balance := bc.GetBalanceUnsafe(tx.From)

	if balance < tx.Amount {
		return false
	}

	return true
}

func (bc *Blockchain) Close() {
	if bc.Storage != nil {
		bc.Storage.Close()
	}
}
