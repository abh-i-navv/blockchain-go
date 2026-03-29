package main

import (
	"blockchain/blockchain"
	"fmt"
)

func main() {
	bc := blockchain.NewBlockchain()

	tx1 := []blockchain.Transaction{
		{From: "David", To: "John", Amount: 10},
		{From: "Tony", To: "Mark", Amount: 10},
	}

	bc.AddBlock(tx1)

	tx2 := []blockchain.Transaction{
		{From: "Carl", To: "Bob", Amount: 10},
	}

	bc.AddBlock(tx2)

	for _, block := range bc.Blocks {
		fmt.Printf("Index: %d\n", block.Index)
		fmt.Printf("Transactions: %+v\n", block.Transactions)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Printf("PrevHash: %s\n", block.PrevHash)
		fmt.Println("---------------------------")
	}

	fmt.Println("Valid Blockchain: ", bc.IsValid())
}
