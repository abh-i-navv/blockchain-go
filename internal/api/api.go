package api

import (
	"blockchain/blockchain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type API struct {
	bc *blockchain.Blockchain
}

type CreateTxRequest struct {
	From   []byte  `json:"from"`
	To     []byte  `json:"to"`
	Amount float64 `json:"amount"`
}

type MineRequest struct {
	Miner []byte `json:"miner"`
}

type BalanceRequest struct {
	Address []byte `json:"address"`
}

func NewAPI(bc *blockchain.Blockchain) *API {
	return &API{bc: bc}
}

func (api *API) GetBlocks(c *gin.Context) {
	c.JSON(http.StatusOK, api.bc.Blocks)
}

func (api *API) CreateTransaction(c *gin.Context) {
	var req CreateTxRequest

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	tx := blockchain.Transaction{
		From:   req.From,
		To:     req.To,
		Amount: req.Amount,
	}

	api.bc.AddBlock([]blockchain.Transaction{tx}, req.From)

	c.JSON(http.StatusOK, gin.H{
		"message": "transaction added",
	})
}

func (api *API) Mine(c *gin.Context) {
	var req MineRequest

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}

	api.bc.AddBlock([]blockchain.Transaction{}, req.Miner)

	c.JSON(200, gin.H{
		"message": "block mined",
	})
}

func (api *API) GetBalance(c *gin.Context) {
	var req BalanceRequest

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	balance := api.bc.GetBalance(req.Address)

	c.JSON(http.StatusOK, gin.H{
		"balance": balance,
	})
}
