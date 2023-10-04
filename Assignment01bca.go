package main

import (
	//"github.com/muqaduszulfiqar /assignment01bca"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
)

type BLOCKCHAIN struct {
	current_hash  string
	previous_hash string
	data          string
}
type chain struct {
	blocks        [7]BLOCKCHAIN
	current_block int
	root_block    int
}

func (b *chain) CalculateHash(stringToHash string) string {
	Hash := md5.Sum([]byte(stringToHash))
	str := hex.EncodeToString(Hash[:])
	return str
}

func (b *chain) NewBlock(transaction string, nonce int) {

	b.blocks[b.current_block].data = transaction
	transaction = transaction + fmt.Sprint(nonce)
	b.blocks[b.current_block].current_hash = b.CalculateHash(transaction)

	if b.root_block != -1 {
		b.blocks[b.current_block].previous_hash = b.blocks[b.current_block-1].current_hash
		b.current_block++
	} else {
		b.root_block++
		b.current_block++
	}

}

func (b *chain) ListBlocks() {
	for i := b.root_block; i < b.current_block; i++ {
		fmt.Println("\nBLOCK NUMBER : ", i+1)

		fmt.Println("\n\n--> THE HASH OF THE CURRENT BLOCK IS  : ", b.blocks[i].current_hash)
		fmt.Println("\n--> THE HASH OF THE PREVIOUS BLOCK IS : ", b.blocks[i].previous_hash)
		fmt.Println("\n--> THE HASH OF THIS BLOCK IS         : ", b.blocks[i].data)
		fmt.Println("\n\n")
	}

}

func (b *chain) VerifyChain() {
	fmt.Println("\nVERIFYING ALL THE BLOCKS")
	for i := 1; i < b.current_block; i++ {
		if b.blocks[i].previous_hash != b.blocks[i-1].current_hash {
			fmt.Println("\n-->INVALID BLOCK CHAIN")
			return
		}
	}

	fmt.Println("\n-->BLOCK CHAIN IS VALID\n\n")

}

func (b *chain) ChangeBlock(index int, data string, nonce int) {
	b.blocks[index].data = data
	data = data + fmt.Sprint(nonce)
	b.blocks[index].current_hash = b.CalculateHash(data)

}

func main() {

	node := new(chain)

	node.current_block = 0
	node.root_block = -1

	node.NewBlock("maryam to hajra ", rand.Intn(100))
	node.NewBlock("umer to usama ", rand.Intn(100))
	node.NewBlock("muqadus to khadija ", rand.Intn(100))
	node.ListBlocks()
	node.VerifyChain()
	node.ChangeBlock(1, "amna to  hussein", 55)
	node.VerifyChain()

}
