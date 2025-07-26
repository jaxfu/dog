package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/jaxfu/dog/pkg/shared"
	"github.com/jessevdk/go-flags"
)

const (
	LINES_PREALLOC_CAP  uint = 1024
	LINES_PREALLOC_SIZE uint = 0
)

func main() {
	// dev logger
	devlog := log.New(os.Stderr, "[DEV] ", log.Ltime)

	// parse cli args
	var cliArgs struct {
		Filepath string `short:"f" long:"filepath" description:"path to target file" required:"true"`
	}
	_, err := flags.Parse(&cliArgs)
	if err != nil {
		// flags.Parse prints err for some reason?
		os.Exit(1)
	}

	// sanitize filepath arg
	safepath, err := sanitizePath(cliArgs.Filepath)
	if err != nil {
		fmt.Printf("invalid filepath '%s':\n%+v\n", cliArgs.Filepath, err)
		os.Exit(1)
	}

	// open file
	file, err := os.Open(safepath)
	if err != nil {
		fmt.Printf("error opening file:\n%+v\n", err)
	}
	defer file.Close()

	// read file into scanner
	scanner := bufio.NewScanner(file)

	// process lines
	lines := make([]shared.Line, LINES_PREALLOC_SIZE, LINES_PREALLOC_CAP)
	var line uint = 1
	for scanner.Scan() {
		lines = append(lines, shared.Line{
			Address: shared.Address{
				Filepath: safepath,
				Linenum:  line,
			},
			Content: scanner.Text(),
		})
		line++
	}
	if scanner.Err() != nil {
		fmt.Printf("error opening file:\n%+v\n", scanner.Err())
		os.Exit(1)
	}

	// TODO: dev
	for _, l := range lines {
		devlog.Printf("%+v\n", l)
	}
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
