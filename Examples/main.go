package main

import (
	"flag"
	"fmt"
	parser "link_parser"
	"os"
)

func main() {
	fileName := flag.String("file", "ex1.html", "file to extract links from")
	flag.Parse()
	file, err := os.Open(*fileName)
	if err != nil {
		fmt.Errorf("failed opening file %q, %d", fileName, err)
	}
	arr, err := parser.Parse(file)
	fmt.Println(arr)

}
