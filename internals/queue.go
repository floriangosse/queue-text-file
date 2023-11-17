package internals

import (
	"os"
	"strings"
)

func ReadFromQueue(filepath string, count int) ([]string, error) {
	inputBytes, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	inputContent := string(inputBytes)
	inputContent = strings.TrimSpace(inputContent)
	inputLines := strings.Split(inputContent, "\n")

	if count > len(inputLines) {
		count = len(inputLines)
	}

	queueLines := inputLines[:count]
	if len(queueLines) == 1 && queueLines[0] == "" {
		queueLines = []string{}
	}

	outputLines := inputLines[count:]
	outputContent := strings.Join(outputLines, "\n")
	outputBytes := []byte(outputContent)

	// Write output to file
	err = os.WriteFile(filepath, outputBytes, 0644)
	if err != nil {
		return nil, err
	}

	return queueLines, nil
}
