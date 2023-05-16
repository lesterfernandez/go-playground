package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		file, err := os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		printLines(file)
		file.Close()
	} else {
		printLines(os.Stdin)
	}
}

func printLines(f *os.File) {
	s := bufio.NewScanner(f)
	for s.Scan() {
		fmt.Println(s.Text())
	}
}
