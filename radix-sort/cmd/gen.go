package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: gen.go [directory] [number]")
		os.Exit(1)
	}

	f, fileErr := os.Open(os.Args[1])
	if fileErr != nil {
		fmt.Printf("Unable to find directory %s\n", os.Args[1])
		os.Exit(1)
	}
	defer f.Close()

	stat, _ := f.Stat()
	if !stat.IsDir() {
		fmt.Println("Usage: gen.go [directory] [number]")
		os.Exit(1)
	}

	// TODO: generate numbers here
}
