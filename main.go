package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func reverse(lines []string) {
	for i := len(lines) - 1; i >= 0; i-- {
		fmt.Println(lines[i])
	}
}

func tac(r io.Reader) {
	scanner := bufio.NewScanner(r)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %s\n", err)
		return
	}
	reverse(lines)
}

func main() {
	if len(os.Args) == 1 {
		tac(os.Stdin)
		return
	}

	for _, filename := range os.Args[1:] {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file %s: %s\n", filename, err)
			continue
		}
		tac(file)
		file.Close()
	}
}
