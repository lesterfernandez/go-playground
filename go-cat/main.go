package main

import (
	"io"
	"log"
	"os"
)

func main() {
	var src *os.File

	if len(os.Args) == 2 {
		file, err := os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}

		src = file
		defer file.Close()
	} else {
		src = os.Stdin
	}

	io.Copy(os.Stdout, src)

	// s := bufio.NewScanner(src)
	// for s.Scan() {
	// 	fmt.Println(s.Text())
	// }
}
