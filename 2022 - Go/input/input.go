package input

import (
	"bufio"
	"log"
	"os"
)

// ReadLines reads a valid AoC input path and
// splits the file into a slice of strings, using
// newline as a delimiter.
func ReadLines(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	log.Printf("open: %s", path)

	lines := make([]string, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		l := scanner.Text()
		lines = append(lines, l)
	}

	return lines, nil
}
