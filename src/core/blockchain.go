package core

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

//type Blockchain struct{
//	Blocks []*Block
//}

const dbFile        = "blockchain.db"
const blocksBucket  = "blocks"

type Blockchain struct{
	tip []byte
	Db *bolt.DB
}

type BlockchainIterator struct{
	currentHash []byte
	Db  *bolt.DB
}

func (bc *Blockchain)AddBlock(data string){
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data,prevBlock.Hash)
	bc.Blocks = append(bc.Blocks,newBlock)
}


func NewBlockchain()*Blockchain{
	//return &Blockchain{[]*Block{NewGenesisBlock()}}

	var tip []byte
	db,err := bolt.Open(dbFile,0600,nil)

	if err != nil{
		log.Panic(err)
	}

	err = db.Update(func(tx *bolt.Tx)error{
		b := tx.Bucket([]byte(blocksBucket))

		if b==nil{
			fmt.Println("No existing blockchain dound.createing a new one")
			genesis := NewGenesisBlock()

			b,err := tx.CreateBucket([]byte(blocksBucket))

			if err !=nil{
				log.Panic(err)
			}

			err = b.Put(genesis.Hash,genesis.Serializee())

			if err != nil{
				log.Panic(err)
			}

			err = b.Put([]byte("l"),genesis.Hash)
			if err != nil{
				log.Panic(err)
			}

			tip = genesis.Hash

		}else{
			tip = b.Get([]byte("l"))
		}
		return nil
	})

	if err != nil{
		log.Panic(err)
	}

	bc := Blockchain{tip,db}
	return &bc
}