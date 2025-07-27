package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jaxfu/dog"
	"github.com/jessevdk/go-flags"
)

func main() {
	// parse cli args
	var cliArgs struct {
		Filepath string `short:"f" long:"filepath" description:"path to target file" required:"true"`
	}
	_, err := flags.Parse(&cliArgs)
	if err != nil {
		// flags.Parse prints err for some reason?
		os.Exit(1)
	}

	lines, err := dog.Get(cliArgs.Filepath, dog.DogOptions{})
	if err != nil {
		fmt.Println(err)
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
