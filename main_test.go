package main

import (
	"bytes"
	"testing"
)

// TestCountWords tests the count function
func TestCountWords(t *testing.T) {
	// Create a buffer of bytes with four words that fits the reader interface
	b := bytes.NewBufferString("word1 word2 word3 word4\n")

	exp := 4

	res := count(b)

	if res != exp {
		t.Errorf("Expected %d, got %d instead. \n", exp, res)
	}
}
