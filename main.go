package main

import (
	"blockchain/blockchain"
	"blockchain/internal/api"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func GetEnv(key, fallback string) string {
	if key == "" {
		return fallback
	}
	return os.Getenv(key)
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("error in .env file")
	}

	bc := blockchain.NewBlockchain()
	defer bc.Close()

	api := api.NewAPI(bc)
	r := gin.Default()

	r.GET("/blocks", api.GetBlocks)
	r.POST("/transation", api.CreateTransaction)
	r.POST("/mine", api.Mine)
	r.POST("/balance", api.GetBalance)

	// http.HandleFunc("/blocks", api.GetBlocks)
	// http.HandleFunc("/transaction", api.CreateTransaction)
	// http.HandleFunc("/mine", api.Mine)
	// http.HandleFunc("/balance", api.GetBalance)

	port := ":" + GetEnv("PORT", "8080")
	fmt.Println("Server running on", port)
	log.Fatal(http.ListenAndServe(port, nil))

	alice := blockchain.NewWallet()
	bob := blockchain.NewWallet()
	carl := blockchain.NewWallet()

	// mining money
	bc.AddBlock([]blockchain.Transaction{}, alice.Public)

	tx1 := blockchain.Transaction{
		From:   alice.Public,
		To:     bob.Public,
		Amount: 40,
	}

	tx1.Sign(alice.Private)

	bc.AddBlock([]blockchain.Transaction{tx1}, []byte{})

	tx2 := blockchain.Transaction{
		From:   alice.Public,
		To:     carl.Public,
		Amount: 40,
	}

	tx1.Sign(alice.Private)

	bc.AddBlock([]blockchain.Transaction{tx2}, []byte{})

	fmt.Println("Alice balance: ", bc.GetBalance(alice.Public))
	fmt.Println("Bob balance: ", bc.GetBalance(bob.Public))
	fmt.Println("Carl balance: ", bc.GetBalance(carl.Public))
}
