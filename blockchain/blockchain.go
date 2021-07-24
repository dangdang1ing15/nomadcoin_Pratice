package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

// #1 데이터를 오직 block에다가만 저장하는 blockchain

type block struct {
	Data     string
	Hash     string
	PrevHash string
}

// #2 blockchain struct 생성해 block들의 slice를 가지고 있다 정의

type blockchain struct {
	Blocks []*block
}

// singleton 패턴
var b *blockchain
var once sync.Once

func (b *block) calculateHash() {
	hash := sha256.Sum256([]byte(b.Data + b.PrevHash))
	b.Hash = fmt.Sprintf("%x", hash)
}

func getLastHash() string {
	totalBlocks := len(GetBlockchain().Blocks)
	if totalBlocks == 0 {
		return ""
	}
	return GetBlockchain().Blocks[totalBlocks-1].Hash
}

func createBlock(data string) *block {
	newBlock := block{data, "", getLastHash()}
	newBlock.calculateHash()
	return &newBlock
}

func (b *blockchain) AddBlock(data string) {
	b.Blocks = append(b.Blocks, createBlock(data))
}

func GetBlockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{}
			b.AddBlock("Genesis")
		})
	}
	return b
}

func (b *blockchain) AllBlocks() []*block {
	return b.Blocks
}