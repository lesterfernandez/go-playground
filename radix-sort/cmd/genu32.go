package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) != 3 {
		handleInvalidUsage()
	}

	dir, dirErr := os.Open(os.Args[1])
	if dirErr != nil {
		handleInvalidUsage()
	}

	stat, _ := dir.Stat()
	if !stat.IsDir() {
		handleInvalidUsage()
	}

	dir.Close()

	n, convErr := strconv.ParseUint(os.Args[2], 10, 0)
	if convErr != nil {
		handleInvalidUsage()
	}

	file, _ := os.Create(dir.Name() + "/nums.txt")
	defer file.Close()

	r := rand.New(rand.NewSource(time.Now().UnixMilli()))
	w := bufio.NewWriter(file)
	defer w.Flush()

	for i := uint64(0); i < n; i++ {
		fmt.Fprintln(w, r.Uint32())
	}
}

func handleInvalidUsage() {
	fmt.Println("Usage: genu32.go [directory] [number]")
	os.Exit(1)
}
