package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("provide input and output filenames")
		os.Exit(-1)
	}

	in := os.Args[1]
	out := os.Args[2]

	data, err := os.ReadFile(in)
	if err != nil {
		fmt.Println("error reading the file")
		log.Fatal(err)
	}
	_, err = os.Stdout.Write(data)
	if err != nil {
		fmt.Println("error wring the data to console")
		return
	}

	err = os.WriteFile(out, data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
