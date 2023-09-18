package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/lesterfernandez/go-playground/radix-sort/radix"
)

func main() {
	if len(os.Args) != 3 {
		handleInvalidUsage()
	}

	n, convErr := strconv.Atoi(os.Args[2])
	if convErr != nil {
		handleInvalidUsage()
	}

	file, err := radix.CreateFile(os.Args[1], "/nums.txt")
	if err != nil {
		handleInvalidUsage()
	}
	defer file.Close()

	r := rand.New(rand.NewSource(time.Now().UnixMilli()))
	w := bufio.NewWriter(file)
	defer w.Flush()

	for i := 0; i < n; i++ {
		fmt.Fprintln(w, r.Uint32())
	}
}

func handleInvalidUsage() {
	fmt.Println("Usage: genu32.go [directory] [number]")
	os.Exit(1)
}
