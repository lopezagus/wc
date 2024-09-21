package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// The count function receives any Reader interface as input which allows us
// to work with different readable inputs such as files or text.
func count(r io.Reader) int {
	// We use a scanner struct instance to read the text fronm the Reader interface
	scanner := bufio.NewScanner(r)

	// Now we define the scanner split type to words (default is lines)
	scanner.Split(bufio.ScanWords)

	// Define the counter variable
	wc := 0

	// Now we scan the input and count words
	for scanner.Scan() {
		wc++
	}

	// Return the total words counted
	return wc
}

func main() {
	// The count function receives the standard input and prints its output
	fmt.Println(count(os.Stdin))
}
