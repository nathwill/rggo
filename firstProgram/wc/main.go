package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	lines := flag.Bool("l", false, "Count lines")
	bytes := flag.Bool("b", false, "Count bytes")
	flag.Parse()

	fmt.Println(count(os.Stdin, *lines, *bytes))
}

func count(r io.Reader, countLines bool, countBytes bool) int {
	scanner := bufio.NewScanner(r)
	if !countLines && !countBytes {
		scanner.Split(bufio.ScanWords)
	}

	wc := 0

	for scanner.Scan() {
		if !countBytes {
			wc++
		} else {
			wc = len(scanner.Bytes())
		}
	}

	return wc
}
