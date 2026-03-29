package blockchain

import (
	"encoding/json"
	"errors"
	"fmt"
	"sort"

	"github.com/dgraph-io/badger"
)

type Storage struct {
	db *badger.DB
}

func NewStorage(path string) (*Storage, error) {
	opts := badger.DefaultOptions(path)
	opts.Logger = nil

	db, err := badger.Open(opts)
	if err != nil {
		return nil, errors.New(fmt.Sprint("error connecting to db", err))
	}

	return &Storage{db: db}, nil
}

func (s *Storage) SaveBlock(block Block) error {
	return s.db.Update(func(txn *badger.Txn) error {
		data, err := json.Marshal(block)
		if err != nil {
			return err
		}

		key := []byte(fmt.Sprintf("block-%d", block.Index))
		return txn.Set(key, data)
	})
}

func (s *Storage) LoadBlocks() ([]Block, error) {
	blocks := []Block{}

	err := s.db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()

		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			val, err := item.ValueCopy(nil)
			if err != nil {
				return err
			}
			var block Block
			err = json.Unmarshal(val, &block)
			if err != nil {
				return err
			}

			blocks = append(blocks, block)
		}
		sort.Slice(blocks, func(i, j int) bool {
			return blocks[i].Index < blocks[j].Index
		})
		return nil
	})

	return blocks, err
}

func (s *Storage) Close() error {
	return s.db.Close()
}
