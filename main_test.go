package main

import (
	"bytes"
	"testing"
)

// TestCountWords tests the count function with default parameters
func TestCountWords(t *testing.T) {
	// Create a buffer of bytes with four words that fits the reader interface
	b := bytes.NewBufferString("word1 word2 word3 word4\n")

	exp := 4

	res := count(b, false)

	if res != exp {
		t.Errorf("Expected %d, got %d instead. \n", exp, res)
	}
}

// TestCountLines test the count function with the -l option set to true
func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3\n line2\n line3 word1")

	exp := 3

	res := count(b, true)

	if res != exp {
		t.Errorf("Expected %d got %d instead. \n", exp, res)
	}
}
