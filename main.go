package main

import (
	"flag"
	"fmt"
	"os"
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
	// lines := flag.String("l", "", "Outputs the number of lines in a file")
	// words := flag.String("w", "", "Outputs the number of words in a file")
	// chars := flag.String("m", "", "Outputs the number of charachters in a file")



	flag.Parse()


	// Check if the filename is provided
	if *bytes == "" {
		fmt.Println("Please provide a filename using the -c flag.")
		return
	}


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
	fmt.Printf("Number of bytes in %s: %d\n", *bytes, fileInfo.Size())

}