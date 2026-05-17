# Blockchain in Go

A simplified blockchain node implementation built in Go to explore blockchain fundamentals, cryptography, transaction validation, mining, and backend system design.

## 🚀 Features

- Proof of Work (PoW) mining
- SHA-256 block hashing
- Blockchain validation
- Transaction system
- ECDSA-based wallets and digital signatures
- Mempool for pending transactions
- Double-spending prevention
- Mining rewards (coinbase transactions)
- SQLite-backed persistence
- REST API using Gin
- Environment-based configuration

---

## 🛠️ Tech Stack

### Backend
- Go
- Gin
- SQLite

### Cryptography
- SHA-256
- ECDSA (P-256)

---

## 📦 API Endpoints

### Get blockchain
```http
GET /blocks
