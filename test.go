package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"testing"
	"time"
)

// This func just generate rnd blocks. It doesn't have any actual data in it, just for making sure our functions work properly.
func randomBlock() GoElapseChainBlock {
	return GoElapseChainBlock{0, time.Now().String(), 0, "", "", ""}
}

// TestCalculateHash checks if the CalculateHash function is working right.
// It should always return the same hash for the same block, so we'll check that.
func TestCalculateHash(t *testing.T) {
	block := randomBlock()

	hash1 := CalculateHash(block)
	hash2 := CalculateHash(block)

	// If the two hashes aren't the same, we've got a problem!
	if hash1 != hash2 {
		t.Errorf("hashes don't match: '%s' vs '%s'", hash1, hash2)
	}
}

// TestSignBlock makes sure the SignBlock function works as expected.
// It should always return the same signature for the same block, so we'll check that.
func TestSignBlock(t *testing.T) {
	block := randomBlock()
	key, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	privateKey = key // set the global private key to our new key

	sig1, _ := SignBlock(block)
	sig2, _ := SignBlock(block)

	// If the two signatures aren't the same, we're in trouble!
	if sig1 != sig2 {
		t.Errorf("signatures don't match: '%s' vs '%s'", sig1, sig2)
	}
}

// TestCreateBlock checks if the CreateBlock function is doing its job.
// It should return a new block with the right index and the right previous hash.
func TestCreateBlock(t *testing.T) {
	oldBlock := GoElapseChainBlock{Index: 0, Hash: "abc123"}
	newBlock, _ := CreateBlock(oldBlock, 456)

	// Check if the new block's index is one more than the old block's index
	if newBlock.Index != oldBlock.Index+1 {
		t.Errorf("wrong block index: got '%d', expected '%d'", newBlock.Index, oldBlock.Index+1)
	}

	// Check if the new block's previous hash is the same as the old block's hash
	if newBlock.PrevHash != oldBlock.Hash {
		t.Errorf("wrong previous hash: got '%s', expected '%s'", newBlock.PrevHash, oldBlock.Hash)
	}
}

