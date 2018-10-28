package core

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

var (
	maxNonce = math.MaxInt64
)

const targetBits = 20

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

//主要是得到一个某个数左移n位后的整数target
func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	fmt.Println("target = ")
	fmt.Println(target)
	target.Lsh(target, uint(256-targetBits))
	fmt.Println("target 左移后 ")
	fmt.Println(target)
	pow := &ProofOfWork{b, target}
	return pow
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			IntToHex(int64(pow.block.Timestamp)),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)
	return data
}

//将区块数据加上n（左移操作中一个数），加上自增数nonce，得到一个大数，转hash，再转整数
//再与原来左移操作后的target比较，比它小就停止循环
func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0
	fmt.Printf("mining the block contains \"%s\" \n", pow.block.Data)
	for nonce < maxNonce {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:]) //转成整数
		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Printf("\n\n")
	return nonce, hash[:]
}
