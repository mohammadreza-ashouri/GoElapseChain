package main

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
	"math/rand"
	"fmt"
)

type Block struct {
	Index     int
	Timestamp string
	Data      int
	Hash      string
	PrevHash  string
}

var Blockchain []Block

func CalculateHash(block Block) string {
	record := string(block.Index) + block.Timestamp + string(block.Data) + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func CreateBlock(oldBlock Block, data int) (Block, error) {
	var newBlock Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.Data = data
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = CalculateHash(newBlock)

	return newBlock, nil
}

func ProofOfElapsedTime() int {
	waitTime := rand.Intn(1000)
	time.Sleep(time.Duration(waitTime) * time.Millisecond)
	return waitTime
}

func main() {
	genesisBlock := Block{0, time.Now().String(), 0, "", ""}
	genesisBlock.Hash = CalculateHash(genesisBlock)
	Blockchain = append(Blockchain, genesisBlock)

	previousBlock := genesisBlock

	for i := 0; i < 10; i++ {
		waitTime := ProofOfElapsedTime()
		newBlock, _ := CreateBlock(previousBlock, waitTime)
		Blockchain = append(Blockchain, newBlock)
		previousBlock = newBlock

		fmt.Printf("Block #%d has been added to the blockchain!\n", newBlock.Index)
		fmt.Printf("Hash: %s\n", newBlock.Hash)
		fmt.Printf("Proof of Elapsed Time (PoET): %dms\n", newBlock.Data)
	}
}
