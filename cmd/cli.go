package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jaxfu/dog"
	"github.com/jaxfu/dog/cmd/internal"
)

func main() {
	fpath, opts, err := internal.GetCliArgs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	lines, err := dog.Get(
		fpath,
		opts,
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// get line num of last line ->
	// convert to int ->
	// convert to string ->
	// get len of string
	// (it's so ugly I love it)
	maxlen := len(strconv.Itoa(int(lines[len(lines)-1].Address.LineNum)))
	for i, l := range lines {
		numStr := fmt.Sprintf("%*d", maxlen, i+1)
		line := fmt.Sprintf("%s| %s", numStr, l.Content)
		fmt.Println(line)
	}
}
