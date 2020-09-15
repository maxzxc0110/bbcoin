package core

import (
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

func (i *BlockchainIterator) Next() *Block {
	var block *Block

	err := i.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodedBlock := b.Get(i.currentHash)
		block = DeserializeBlock(encodedBlock)

		return nil
	})
	if err != nil{
		log.Panic(err)
	}

	i.currentHash = block.PrevBlockHash

	return block
}

func (bc *Blockchain) Iterator() *BlockchainIterator {
	bci := &BlockchainIterator{bc.tip, bc.Db}

	return bci
}



func (bc *Blockchain)AddBlock(data string){

	//var lastHash []byte
	//
	//err := bc.Db.View(func(tx *bolt.Tx) error {
	//	b := tx.Bucket([]byte(blocksBucket))
	//	lastHash = b.Get([]byte("l"))
	//
	//	return nil
	//})
	//if err != nil{
	//	log.Panic(err)
	//}




	//newBlock := NewBlock(data, lastHash)

	//err := bc.Db.Update(func(tx *bolt.Tx) error {
	//
	//	b := tx.Bucket([]byte(blocksBucket))
	//
	//	if b != nil{
	//		byteBytes := b.Get(bc.tip)
	//
	//		block := DeserializeBlock(byteBytes)
	//
	//		newBlock := NewBlock(data,block.Hash)
	//
	//		err := b.Put(newBlock.Hash, newBlock.Serialize())
	//		if err != nil{
	//			log.Panic(err)
	//		}
	//		err = b.Put([]byte("l"), newBlock.Hash)
	//		if err != nil{
	//			log.Panic(err)
	//		}
	//		bc.tip = newBlock.Hash
	//	}
	//
	//	return nil
	//})
	//if err != nil{
	//	log.Panic(err)
	//}

	err := bc.Db.Update(func(tx *bolt.Tx) error {
		//1.获取表
		b := tx.Bucket([]byte(blocksBucket))
		//2.创建新区块
		if b != nil {
			//通过Key:blc.Tip获取Value(区块序列化字节数组)
			byteBytes := b.Get(bc.tip)
			//反序列化出最新区块(上一个区块)对象
			block := DeserializeBlock(byteBytes)

			//3.通过NewBlock进行挖矿生成新区块newBlock
			newBlock := NewBlock(data, block.Hash)
			//4.将最新区块序列化并且存储到数据库中(key=新区块的Hash值，value=新区块序列化)
			err := b.Put(newBlock.Hash, newBlock.Serialize())
			if err != nil {
				log.Panic(err)
			}

			/*5.更新数据库中"l"对应的Hash为新区块的Hash值
			用途:便于通过该Hash值找到对应的Block序列化，从而找到上一个Block对象，为生成新区块函数NewBlock提供高度Height与上一个区块的Hash值PreBlockHash
			*/
			err = b.Put([]byte("l"), newBlock.Hash)
			if err != nil {
				log.Panic(err)
			}
			//6. 更新Tip值为新区块的Hash值
			bc.tip = newBlock.Hash
		}

		return nil

	})
	if err != nil {
		log.Panic(err)
	}
}


func NewBlockchain()*Blockchain{
	//return &Blockchain{[]*Block{NewGenesisBlock()}}

	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))

		if b == nil {
			genesis := NewGenesisBlock()
			b, err := tx.CreateBucket([]byte(blocksBucket))
			if err != nil{
				log.Panic(err)
			}
			err = b.Put(genesis.Hash, genesis.Serialize())
			if err != nil{
				log.Panic(err)
			}
			err = b.Put([]byte("l"), genesis.Hash)
			if err != nil{
				log.Panic(err)
			}
			tip = genesis.Hash
		} else {
			tip = b.Get([]byte("l"))
		}

		return nil
	})
	if err != nil{
		log.Panic(err)
	}

	bc := Blockchain{tip, db}

	return &bc
}