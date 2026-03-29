package api

import (
	"blockchain/blockchain"
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
)

type API struct {
	bc *blockchain.Blockchain
}

type CreateTxRequest struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
}

type MineRequest struct {
	Miner string `json:"miner"`
}

type BalanceRequest struct {
	Address string `json:"address"`
}

func NewAPI(bc *blockchain.Blockchain) *API {
	return &API{bc: bc}
}

func parseAddress(value string) []byte {
	decoded, err := base64.StdEncoding.DecodeString(value)
	if err == nil {
		return decoded
	}

	return []byte(value)
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
		From:   parseAddress(req.From),
		To:     parseAddress(req.To),
		Amount: req.Amount,
	}

	api.bc.AddBlock([]blockchain.Transaction{tx}, parseAddress(req.From))

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
		return
	}

	api.bc.AddBlock([]blockchain.Transaction{}, parseAddress(req.Miner))

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

	balance := api.bc.GetBalance(parseAddress(req.Address))

	c.JSON(http.StatusOK, gin.H{
		"balance": balance,
	})
}
