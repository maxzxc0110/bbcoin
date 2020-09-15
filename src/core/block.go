package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
	"strconv"
	"time"
)

type Block struct {
	Timestamp int64
	Data 			[]byte
	PrevBlockHash 	[]byte
	Hash 			[]byte
	Nonce	int
}

func NewBlock(data string,PrevBlockHash []byte)*Block{
	block := &Block{time.Now().Unix(),[]byte(data),PrevBlockHash,[]byte{},0}
	//block.SetHash()
	pow := NewProofOfWork(block)
	noce,hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = noce

	return block
}

func (b *Block) SetHash(){
	timestamp := []byte(strconv.FormatInt(b.Timestamp,10))
	headers := bytes.Join([][]byte{b.PrevBlockHash,b.Data,timestamp},[]byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}


func (b *Block)Serialize()[]byte{
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil{
		log.Panic(err)
	}

	return result.Bytes()
}

func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)

	if err!=nil{
		fmt.Println(err)
	}

	return &block
}

func NewGenesisBlock()*Block{
	return NewBlock("Genesis Block",[]byte{})
}