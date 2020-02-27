package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/otiai10/gosseract"
)

const (
	dataPathPrefix = "./data/"
)

func main() {
	fmt.Println("start.")

	// Init OCR client.
	client := gosseract.NewClient()
	defer client.Close()

	files, err := ioutil.ReadDir(dataPathPrefix)
	if err != nil {
		fmt.Println("Failed to get files: ", err)
		os.Exit(1)
	}

	for _, fileName := range files {
		// Set data.
		client.SetImage(dataPathPrefix + fileName.Name())

		// Get the answer from data.
		text, err := client.Text()
		if err != nil {
			fmt.Println("Failed to get answer: ", err)
			os.Exit(1)
		}

		// Show the answer.
		fmt.Printf("Got answer from %s: %s\n", fileName.Name(), text)
	}
}
