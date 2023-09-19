package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/lesterfernandez/go-playground/radix-sort/radix"
)

func main() {
	if len(os.Args) != 2 {
		handleInvalidUsage()
	}

	f, e := fileFromArgs(os.Args[1])
	if e != nil {
		handleInvalidUsage()
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	nums := readNums(scanner)

	radix.BinaryRadixSort(nums, 16)
	writeResult(nums)
}

func fileFromArgs(filename string) (*os.File, error) {
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

func readNums(s *bufio.Scanner) []uint {
	nums := make([]uint, 0)
	for s.Scan() {
		line := s.Text()
		num, parseErr := strconv.ParseUint(line, 10, 32)
		if parseErr != nil {
			panic(parseErr)
		}
		nums = append(nums, uint(num))
	}
	return nums
}

func handleInvalidUsage() {
	fmt.Println("Usage: u32.go [filename]")
	os.Exit(1)
}

func writeResult(nums []uint) {
	output, _ := os.Create("result.txt")
	w := bufio.NewWriter(output)
	defer w.Flush()
	for _, num := range nums {
		fmt.Fprintln(w, num)
	}
}
