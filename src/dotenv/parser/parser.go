package parser

import (
	"bufio"
	"io"
)

func Parse(r io.Reader) ([]string, error) {
	lines := make([]string, 0, 10)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}
