package main

import (
	"blockchain/blockchain"
	"blockchain/internal/api"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

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

	api := api.NewAPI(bc)
	r := gin.Default()

	r.GET("/blocks", api.GetBlocks)
	r.POST("/transaction", api.CreateTransaction)
	r.POST("/mine", api.Mine)
	r.POST("/balance", api.GetBalance)
	r.POST("/wallet", func(c *gin.Context) {
		wallet := blockchain.NewWallet()

		c.JSON(200, gin.H{
			"public": wallet.Public,
		})
	})

	// http.HandleFunc("/blocks", api.GetBlocks)
	// http.HandleFunc("/transaction", api.CreateTransaction)
	// http.HandleFunc("/mine", api.Mine)
	// http.HandleFunc("/balance", api.GetBalance)

	port := ":" + GetEnv("PORT", "8080")
	fmt.Println("Server running on", port)

	srv := http.Server{
		Addr:    port,
		Handler: r,
	}

	alice := blockchain.NewWallet()
	bob := blockchain.NewWallet()
	carl := blockchain.NewWallet()

	// mining money
	// bc.AddBlock([]blockchain.Transaction{}, alice.Public)

	// tx1 := blockchain.Transaction{
	// 	From:   alice.Public,
	// 	To:     bob.Public,
	// 	Amount: 40,
	// }

	// tx1.Sign(alice.Private)

	// bc.AddBlock([]blockchain.Transaction{tx1}, []byte{})

	// tx2 := blockchain.Transaction{
	// 	From:   alice.Public,
	// 	To:     carl.Public,
	// 	Amount: 40,
	// }

	// tx1.Sign(alice.Private)

	// bc.AddBlock([]blockchain.Transaction{tx2}, []byte{})

	fmt.Println("Alice balance: ", bc.GetBalance(alice.Public))
	fmt.Println("Bob balance: ", bc.GetBalance(bob.Public))
	fmt.Println("Carl balance: ", bc.GetBalance(carl.Public))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-quit
		fmt.Println("Shutting down...")

		bc.Close()
		os.Exit(0)
	}()

	log.Fatal(srv.ListenAndServe())
}
