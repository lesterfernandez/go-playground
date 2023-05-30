package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	var url string

	if len(os.Args) != 2 {
		fmt.Scanln(&url)
	} else {
		url = os.Args[1]
	}

	resp, err := http.Get(url)

	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	doc, err := html.Parse(bytes.NewReader(body))
	if err != nil {
		log.Fatalln(err)
	}

	words, pics := 0, 0
	count(doc, &words, &pics)
	fmt.Printf("Counted %d words and %d pics\n", words, pics)
}

func count(n *html.Node, words, pics *int) {
	if n.Type == html.TextNode {
		*words += len(strings.Fields(n.Data))
	} else if n.Type == html.ElementNode && n.Data == "img" {
		*pics++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		count(c, words, pics)
	}
}
