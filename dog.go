package dog

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

const (
	LINES_PREALLOC_CAP  uint = 1024
	LINES_PREALLOC_SIZE uint = 0
)

type DogOptions struct{}

type Line struct {
	Address Address `json:"address"`
	Content string  `json:"content"`
}

type Address struct {
	Filepath string `json:"filepath"`
	LineNum  uint   `json:"line_num"`
}

func Get(filepath string, opts DogOptions) ([]Line, error) {
	// sanitize filepath
	fpath, err := sanitizePath(filepath)
	if err != nil {
		return nil, fmt.Errorf("invalid filepath '%s':\n%+v\n", filepath, err)
	}

	// open file
	file, err := os.Open(fpath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	// wrap file in scanner
	scanner := bufio.NewScanner(file)

	lines, err := process(scanner, fpath)
	if err != nil {
		return nil, fmt.Errorf("error processing file '%s':\n%+v\n", filepath, err)
	}

	return lines, nil
}

func sanitizePath(fpath string) (string, error) {
	if filepath.IsAbs(fpath) {
		return fpath, nil
	}
	absPath, err := filepath.Abs(fpath)
	if err != nil {
		return "", err
	}

	return absPath, nil
}

func process(scanner *bufio.Scanner, fpath string) ([]Line, error) {
	// process lines
	lines := make([]Line, LINES_PREALLOC_SIZE, LINES_PREALLOC_CAP)
	var line uint = 1
	for scanner.Scan() {
		lines = append(lines, Line{
			Address: Address{
				Filepath: fpath,
				LineNum:  line,
			},
			Content: scanner.Text(),
		})
		line++
	}
	if scanner.Err() != nil {
		return nil, fmt.Errorf("error opening file:\n%+v\n", scanner.Err())
	}

	return lines, nil
}
