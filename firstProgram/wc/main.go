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
	fList := flag.Args()

	if len(fList) > 0 {
		for _, f := range fList {
			r, err := os.Open(f)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			fmt.Println(f, count(r, *lines, *bytes))
		}
	} else {
		fmt.Println(count(os.Stdin, *lines, *bytes))
	}
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
