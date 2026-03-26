package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type Block struct {
	Index     int
	Timestamp int64
	Data      string
	PrevHash  string
	Hash      string
	Nonce     int
}

func (b *Block) CalculateHash() string {
	record := fmt.Sprintf("%d%d%s%s%d",
		b.Index,
		b.Timestamp,
		b.Data,
		b.PrevHash,
		b.Nonce)

	hash := sha256.Sum256([]byte(record))
	return hex.EncodeToString(hash[:])
}

func NewBlock(index int, data string, prevHash string) Block {
	block := Block{
		Index:     index,
		Timestamp: time.Now().Unix(),
		Data:      data,
		PrevHash:  prevHash,
		Nonce:     0,
	}

	block.Hash = block.CalculateHash()
	return block
}
