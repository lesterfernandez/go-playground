package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f := fileFromArgs()
	defer f.Close()

	scanner := bufio.NewScanner(f)
	nums := readNums(scanner)

	nums = radixSort(nums)
	writeResult(nums)
}

func fileFromArgs() *os.File {
	if len(os.Args) != 2 {
		handleInvalidUsage()
	}

	f, fErr := os.Open(os.Args[1])
	if fErr != nil {
		handleInvalidUsage()
	}

	st, sErr := f.Stat()
	if sErr != nil || st.IsDir() {
		handleInvalidUsage()
	}

	return f
}

func readNums(s *bufio.Scanner) []uint32 {
	nums := make([]uint32, 0)
	for s.Scan() {
		line := s.Text()
		num, parseErr := strconv.ParseUint(line, 10, 32)
		if parseErr != nil {
			panic(parseErr)
		}
		nums = append(nums, uint32(num))
	}
	return nums
}

func handleInvalidUsage() {
	fmt.Println("Usage: main.go [filename]")
	os.Exit(1)
}

func radixSort(nums []uint32) []uint32 {
	fmt.Println(len(nums))
	return nums
}

func writeResult(nums []uint32) {
	output, _ := os.Create("result.txt")
	w := bufio.NewWriter(output)
	defer w.Flush()
	for num := range nums {
		fmt.Fprintln(w, num)
	}
}
