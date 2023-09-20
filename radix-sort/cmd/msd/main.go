package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/lesterfernandez/go-playground/radix-sort/radix"
)

func main() {
	if len(os.Args) != 2 {
		handleInvalidUsage()
	}

	f, e := radix.OpenFile(os.Args[1])
	if e != nil {
		handleInvalidUsage()
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	arr := make([]string, 0)
	for scanner.Scan() {
		arr = append(arr, scanner.Text())
	}

	radix.MsdRadixSort(arr)
	writeResult(arr)
}

func handleInvalidUsage() {
	fmt.Fprintln(os.Stderr, "Usage: hex.go [filename]")
	os.Exit(1)
}

func writeResult(nums []string) {
	output, _ := os.Create("result.txt")
	w := bufio.NewWriter(output)
	defer w.Flush()
	for _, num := range nums {
		fmt.Fprintln(w, num)
	}
}
