package core

import "fmt"

type CLI struct {
	Bc *Blockchain
}

func (cli *CLI)printUsage(){
	fmt.Println("Usage")
	fmt.Println("addblock -data BLOCK_DATA - add a block to the blockchain")
	fmt.Println("printchain -print all the blocks of the blockchain")
}


func (cli *)