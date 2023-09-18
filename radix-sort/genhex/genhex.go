package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
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

	file, err := radix.CreateFile(os.Args[1], "/hex.txt")
	if err != nil {
		handleInvalidUsage()
	}
	defer file.Close()

	r := rand.New(rand.NewSource(time.Now().UnixMilli()))
	w := bufio.NewWriter(file)
	defer w.Flush()

	for i := 0; i < n; i++ {
		b := make([]byte, r.Intn(4)+1)
		r.Read(b)
		fmt.Fprintln(w, strings.ToUpper(hex.EncodeToString(b)))
	}
}

func handleInvalidUsage() {
	fmt.Println("Usage: genhex.go [directory] [number]")
	os.Exit(1)
}
