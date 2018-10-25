package main

import (
	"core"
	"fmt"
)

func main() {
	bc := core.NewBlockChain()
	bc.AddBlock("a send 1 btc to b")
	bc.AddBlock("b earn 20 btc")
	for _, block := range bc.Blocks {
		fmt.Printf("prevHash:%d\n", block.PrevBlockHash)
		fmt.Printf("data:%s\n", block.Data)
		fmt.Printf("hash:%d\n", block.Hash)
	}
}
