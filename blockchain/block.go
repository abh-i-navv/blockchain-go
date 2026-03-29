package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Block struct {
	Index        int
	Timestamp    int64
	Transactions []Transaction
	PrevHash     string
	Hash         string
	Nonce        int
	Difficulty   int
}

func (b *Block) CalculateHash() string {
	txBytes, _ := json.Marshal(b.Transactions)

	record := fmt.Sprintf("%d%d%s%s%d%d",
		b.Index,
		b.Timestamp,
		txBytes,
		b.PrevHash,
		b.Nonce,
		b.Difficulty,
	)

	hash := sha256.Sum256([]byte(record))
	return hex.EncodeToString(hash[:])
}

func NewBlock(index int, txs []Transaction, prevHash string, difficulty int) Block {
	block := Block{
		Index:        index,
		Timestamp:    time.Now().Unix(),
		Transactions: txs,
		PrevHash:     prevHash,
		Nonce:        0,
		Difficulty:   difficulty,
	}

	block.Mine()
	return block
}

func (b *Block) Mine() {
	target := strings.Repeat("0", b.Difficulty)

	for {
		hash := b.CalculateHash()
		if strings.HasPrefix(hash, target) {
			b.Hash = hash
			return
		}
		b.Nonce++
	}
}
