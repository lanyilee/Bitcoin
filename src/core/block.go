package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
	"strconv"
	"time"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run() //挖矿，计算数值
	block.Hash = hash[:]
	block.Nonce = nonce
	//block.SetHash()
	return block
}

func (block *Block) SetHash() {
	time := []byte(strconv.FormatInt(block.Timestamp, 10))                            //int64转byte，先转成string再转byte
	hearders := bytes.Join([][]byte{block.PrevBlockHash, block.Data, time}, []byte{}) //将几个数组合成一个
	hash := sha256.Sum256(hearders)
	block.Hash = hash[:]
}

//创世纪块
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

//序列化
func (block *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(block)
	if err != nil {
		log.Panic(err)
	}
	return result.Bytes()
}

//反序列化
func DeserializedBlock(b []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(b))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}
	return &block
}
