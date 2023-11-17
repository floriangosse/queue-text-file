package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"floriangosse.com/queue-text-file/internals"
)

func main() {
	if len(os.Args) != 3 {
		printError(errors.New("invalid count of arguments"))
	}

	filepath := os.Args[1]
	countAsStr := os.Args[2]

	// Convert count to int
	count, err := strconv.Atoi(countAsStr)
	if err != nil {
		printError(err)
	}

	queueLines, err := internals.ReadFromQueue(filepath, count)
	if err != nil {
		printError(err)
	}

	for _, line := range queueLines {
		fmt.Println(line)
	}
}

func printError(err error) {
	fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
	os.Exit(1)
}
