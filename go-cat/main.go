package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var scanner *bufio.Scanner

	if len(os.Args) == 2 {
		file, err := os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		scanner = bufio.NewScanner(file)
		printFile(scanner)
		file.Close()
	} else {
		scanner = bufio.NewScanner(os.Stdin)
		printFile(scanner)
	}
}

func printFile(s *bufio.Scanner) {
	for s.Scan() {
		fmt.Println(s.Text())
	}
}
