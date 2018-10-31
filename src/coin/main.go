package main

import "core"

func main() {
	bc := core.NewBlockChain()
	defer bc.Db.Close()

	//	bc := core.NewBlockChain()
	//	bc.AddBlock("a send 1 btc to b")
	//	bc.AddBlock("b earn 20 btc")
	//	for _, block := range bc.Blocks {
	//		fmt.Printf("prevHash:%d\n", block.PrevBlockHash)
	//		fmt.Printf("data:%s\n", block.Data)
	//		fmt.Printf("hash:%d\n", block.Hash)
	//		//工作量证明
	//		pow := core.NewProofOfWork(block)
	//		flag := pow.Validate()
	//		fmt.Println("validate:" + strconv.FormatBool(flag))
	//	}
}
