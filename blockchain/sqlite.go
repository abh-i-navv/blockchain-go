package blockchain

import (
	"database/sql"
	"encoding/json"

	_ "modernc.org/sqlite"
)

type SQLiteStorage struct {
	db *sql.DB
}

func NewSQLiteStorage(path string) (*SQLiteStorage, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	query := `
	CREATE TABLE IF NOT EXISTS blocks (
		id INTEGER PRIMARY KEY,
		data BLOB
	);
	`

	_, err = db.Exec(query)

	if err != nil {
		return nil, err
	}

	return &SQLiteStorage{db: db}, nil
}

func (s *SQLiteStorage) SaveBlock(block Block) error {
	data, err := json.Marshal(block)
	if err != nil {
		return nil
	}

	_, err = s.db.Exec(
		"INSERT INTO blocks (id,data) VALUES(?, ?)",
		block.Index,
		data,
	)
	return err
}

func (s *SQLiteStorage) LoadBlocks() ([]Block, error) {
	rows, err := s.db.Query("SELECT data FROM blocks ORDER BY id")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var blocks []Block

	for rows.Next() {
		var data []byte

		if err := rows.Scan(&data); err != nil {
			return nil, err
		}

		var block Block
		if err := json.Unmarshal(data, &block); err != nil {
			return nil, err
		}

		blocks = append(blocks, block)

	}
	return blocks, nil
}

func (s *SQLiteStorage) Close() error {
	return s.db.Close()
}
