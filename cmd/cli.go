package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jaxfu/dog"
	"github.com/jaxfu/dog/cmd/internal"
)

func main() {
	cliArgs, err := internal.GetCliArgs()
	fmt.Printf("%+v\n", cliArgs)
	if err != nil {
		fmt.Printf("error getting cli arguments\n%+v\n", err)
		os.Exit(1)
	}

	lines, err := dog.Get(
		cliArgs.Filepath,
		dog.DogOptions{},
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
