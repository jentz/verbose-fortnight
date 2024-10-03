package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/google/uuid"
)

func main() {

	// get file in and file out from flags args
	fileIn := ""
	flag.StringVar(&fileIn, "in", "-", "input file")
	fileOut := ""
	flag.StringVar(&fileOut, "out", "", "output file")

	flag.Parse()

	// Open the input file
	inputFile, err := os.Open(fileIn)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer inputFile.Close()

	// Create the output file
	outputFile, err := os.Create(fileOut)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer outputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	writer := bufio.NewWriter(outputFile)

	for scanner.Scan() {
		line := scanner.Text()
		var logEntry map[string]interface{}

		// Parse the JSON log entry
		if err := json.Unmarshal([]byte(line), &logEntry); err != nil {
			fmt.Println("Error parsing JSON:", err)
			continue
		}

		// Add the eventId attribute
		logEntry["eventId"] = uuid.New().String()

		// Convert the log entry back to JSON
		modifiedLine, err := json.Marshal(logEntry)
		if err != nil {
			fmt.Println("Error marshalling JSON:", err)
			continue
		}

		// Write the modified line to the output file
		if _, err := writer.WriteString(string(modifiedLine) + "\n"); err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	// Flush the writer to ensure all data is written to the file
	if err := writer.Flush(); err != nil {
		fmt.Println("Error flushing writer:", err)
	}
}
