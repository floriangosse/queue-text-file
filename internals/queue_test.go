package internals

import (
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestReadFromQueue(t *testing.T) {
	// 1. Create queue file that contains 14 lines - each one contains just the line number
	queueFile, err := createQueueFile()
	if err != nil {
		t.Fatal(err)
	}

	queueFilePath := queueFile.Name()

	// 2. Call ReadFromQueue with count = 10
	queueLines, err := ReadFromQueue(queueFilePath, 10)
	if err != nil {
		t.Fatal(err)
	}

	// 3. Check if the first 10 lines were removed from queue file
	testQueueLines(t, queueLines, 10, 0)
	testQueueFile(t, queueFile, 4, 10)

	// 4. Call ReadFromQueue with count = 10
	queueLines, err = ReadFromQueue(queueFilePath, 10)
	if err != nil {
		t.Fatal(err)
	}

	// 5. Check if the last 4 lines were removed from queue file
	testQueueLines(t, queueLines, 4, 10)
	testQueueFile(t, queueFile, 0, 14)
}

func testQueueFile(t *testing.T, queueFile *os.File, expectedCount int, offset int) {
	queueFileBytes, err := os.ReadFile(queueFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	queueFileContent := strings.TrimSpace(string(queueFileBytes))
	queueFileLines := strings.Split(queueFileContent, "\n")
	if len(queueFileLines) == 1 && queueFileLines[0] == "" {
		queueFileLines = []string{}
	}

	testQueueLines(t, queueFileLines, expectedCount, offset)
}

func testQueueLines(t *testing.T, queueLines []string, expectedCount int, offset int) {
	if len(queueLines) != expectedCount {
		t.Fatalf("Expected %d lines but got %d", expectedCount, len(queueLines))
	} else {
		for i := 0; i < expectedCount; i++ {
			expectedLine := strconv.Itoa(i + 1 + offset)
			if queueLines[i] != expectedLine {
				t.Fatalf("Expected line %d to be %s but got %s", i+1, expectedLine, queueLines[i])
			}
		}
	}
}

func createQueueFile() (*os.File, error) {
	queueFile, err := os.CreateTemp("", "*")
	if err != nil {
		return nil, err
	}

	queueFileContent := ""
	for i := 1; i <= 14; i++ {
		queueFileContent += strconv.Itoa(i) + "\n"
	}

	err = os.WriteFile(queueFile.Name(), []byte(queueFileContent), 0644)
	if err != nil {
		return nil, err
	}

	return queueFile, nil
}
