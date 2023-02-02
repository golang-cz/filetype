package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	matcher := New()
	//fmt.Println(matcher)

	if len(os.Args) != 2 {
		fmt.Println(matcher)
		os.Exit(0)
		//log.Fatal("usage: filetype <file>")
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	limitedReader := &io.LimitedReader{R: f, N: 512}
	b, err := io.ReadAll(limitedReader)
	if err != nil {
		log.Fatal(err)
	}

	//spew.Dump(b)

	header := key(b)
	fmt.Println(matcher.MatchString(header))

}
