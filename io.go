package advent20201218

import (
	"bufio"
	"os"
)

// RecordsFromFile gets the lines from a file
func RecordsFromFile(filename string) (lines []string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return
}
