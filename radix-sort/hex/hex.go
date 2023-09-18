package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/lesterfernandez/go-playground/radix-sort/radix"
)

func main() {
	if len(os.Args) != 2 {
		handleInvalidUsage()
	}

	f, e := openFile(os.Args[1])
	if e != nil {
		handleInvalidUsage()
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	arr := make([]string, 0)
	for scanner.Scan() {
		arr = append(arr, scanner.Text())
	}

	arr = radix.MsdRadixSort(arr)
	writeResult(arr)
}

func openFile(filename string) (*os.File, error) {
	f, fErr := os.Open(filename)
	if fErr != nil {
		return nil, errors.New("")
	}

	st, sErr := f.Stat()
	if sErr != nil || st.IsDir() {
		return nil, errors.New("")
	}

	return f, nil
}

func handleInvalidUsage() {
	fmt.Println("Usage: hex.go [filename]")
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
