package main

import (
	"bitcoin/src/core"
)

func main(){
	bc := core.NewBlockchain()

	defer bc.Db.Close()

	cli := core.CLI{bc}
	cli.Run()

	//bc.AddBlock("Send 1 BTC to Ivan")
	//bc.AddBlock("Send 2 more Btc to Ivan")
	//
	//for _,block := range bc.Blocks{
	//	fmt.Printf("preev.hash:%x\n",block.PrevBlockHash)
	//	fmt.Printf("Data:%s\n",block.Data)
	//	fmt.Printf("Hash:%x\n",block.Hash)
	//
	//
	//	pow := core.NewProofOfWork(block)
	//	fmt.Printf("Pow :%s\n",strconv.FormatBool(pow.Validate()))
	//	fmt.Println()
	//}
}