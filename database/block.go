package database

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

type Hash [32]byte

func (h Hash) MarshalText() ([]byte, error) {
	return []byte(hex.EncodeToString(h[:])),nil
}

func (h *Hash) UnmarshalText(data []byte) error {
	_, err := hex.Decode(h[:],data)
	return err
}

type Block struct {
	Header BlockHeader 
	TXs []Tx
}

type BlockHeader struct {
	Parent Hash
	Number uint64
	Time uint64
}

type BlockFS struct {
	Key Hash
	Value Block
}

func NewBlock(parent Hash,number uint64, time uint64, txs []Tx) Block {
	return Block{BlockHeader{parent,number,time},txs}
}

func (b Block) Hash() (Hash, error) {
	blockjson, err := json.Marshal(b)
	if err != nil {
		return Hash{}, err
	}

	return sha256.Sum256(blockjson), nil
}