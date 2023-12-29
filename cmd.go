package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)



func CountChars() int {
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
	// A scanner is used to read text from a Reader (such as files)
	scanner := bufio.NewScanner(in)

	
	scanner.Split(bufio.ScanRunes)


	// Defining a counter
	wc := 0

	// For every word scanned, increment the counter
	for scanner.Scan() {
		wc++
	}

	// return the total
	return wc
}


func CountLines() int {
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
	scanner := bufio.NewScanner(in)
	// Count the number of lines
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	return lineCount
}


func CountWords() int {	
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
	scanner := bufio.NewScanner(in)


	// Count the number of words
	wordCount := 0
	for scanner.Scan() {
		// Split each line into words using strings.Fields
		words := strings.Fields(scanner.Text())
		wordCount += len(words)
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	

	return wordCount
}


func CountBytes() int {
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

	
	var totalBytes int

	// Allocate a buffer with the file size
	buffer := make([]byte, 1024)

	for {
		// Read a chunk from the file
		numBytes, err := in.Read(buffer)
		if err != nil && err != io.EOF {
			fmt.Println("Error reading file:", err)
			os.Exit(1)
		}

		// Break if there is no more content to read
		if numBytes == 0 {
			break
		}
	

		// Increment the total number of bytes read
		totalBytes += numBytes
	}

	return totalBytes
}

