package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)


func main() {

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(),
		"%s tool. Developed for The Coding Challenges\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "Copyright 2023\n")
		fmt.Fprintln(flag.CommandLine.Output(), "Usage information:")
		flag.PrintDefaults()
		}

	// Parsing command line flags
	bytes := flag.String("c", "", "Outputs the number of bytes in a file")
	lines := flag.String("l", "", "Outputs the number of lines in a file")
	words := flag.String("w", "", "Outputs the number of words in a file")
	chars := flag.String("m", "", "Outputs the number of charachters in a file")



	flag.Parse()

	switch {
	case *bytes != "":
		// Open the file
		file, err := os.Open(*bytes)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()

		// Get file information
		fileInfo, err := file.Stat()
		if err != nil {
			fmt.Println("Error getting file information:", err)
			return
		}

		// Print the number of bytes in the file
		fmt.Printf("%d %s\n", fileInfo.Size(), *bytes)
			
	case *lines != "":
		// Open the file
		file, err := os.Open(*lines)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()

		// Create a scanner to read the file line by line
		scanner := bufio.NewScanner(file)

		// Count the number of lines
		lineCount := 0
		for scanner.Scan() {
			lineCount++
		}

		// Check for scanner errors
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading file:", err)
			return
		}

		// Print the number of lines in the file
		fmt.Printf("%d %s\n", lineCount, *lines)

	
	case *words != "":
		// Open the file
		file, err := os.Open(*words)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()

		// Create a scanner to read the file line by line
		scanner := bufio.NewScanner(file)

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
			return
		}

		// Print the number of words in the file
		fmt.Printf("%d %s\n", wordCount, *words)


	case *chars != "":
		// Open the file
		file, err := os.Open(*chars)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()

		// Create a buffer to read the file in chunks
		buffer := make([]byte, 1024)

		// Count the number of characters
		charCount := 0
		for {
			// Read a chunk from the file
			n, err := file.Read(buffer)
			if err != nil && n == 0 {
				break
			}

			// Count the number of characters in the chunk
			charCount += utf8.RuneCount(buffer[:n])
		}

		// Print the number of characters in the file
		fmt.Printf("%d %s\n", charCount, *chars)
		
	
	}




}