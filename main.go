package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

// The count function receives any Reader interface as input which allows us
// to work with different readable inputs such as files or text.
func count(r io.Reader, countLines bool) int {
	// We use a scanner struct instance to read the text from the Reader interface
	scanner := bufio.NewScanner(r)

	// Now we define the scanner split type based on the flag parameter
	if !countLines {
		scanner.Split(bufio.ScanWords)
	}

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
	// Define boolean flag -l parameter to read lines instead of words, default is false
	lines := flag.Bool("l", false, "Count lines")

	// Parse the flags provided by the user
	flag.Parse()

	// The count function receives the standard input and prints its output
	// Pass the command line flags as arguments
	fmt.Println(count(os.Stdin, *lines))
}
