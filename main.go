package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	lines := flag.Bool("l", false, "output number of lines in a file")
	chars := flag.Bool("m", false, "output number of charchters in a file")
	words := flag.Bool("w", false, "output number of words in a file")
	bytes := flag.Bool("c", false, "output number of bytes in a file")

	flag.Parse()



	var in io.Reader
	var filename string

	if filename = flag.Arg(0); filename != "" {
		f, err := os.Open(filename)
		if err != nil {
			fmt.Println("error opening file: err:", err)
			os.Exit(1)
		}
		defer f.Close()

		in = f
	} else {
		in = os.Stdin
	}


	var lineCount, charCount, wordCount, byteCount int


	switch {
		case *lines:
			lineCount = CountLines(in)
			// Print the number of lines in the file
			fmt.Printf("%d %s\n", lineCount, filename)
		case *chars:
			charCount = CountChars(in)
			// Print the number of chars/runes in the file
			fmt.Printf("%d %s\n", charCount, filename)
		case *words:
			wordCount = CountWords(in)
			// Print the number of words in the file
			fmt.Printf("%d %s\n", wordCount, filename)
		case *bytes:
			byteCount = CountBytes(in)
			// Print the number of words in the file
			fmt.Printf("%d %s\n", byteCount, filename)


	}


	
}