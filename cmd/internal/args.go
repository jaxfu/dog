package internal

import (
	"flag"
	"fmt"

	"github.com/jaxfu/dog"
)

type cliArgs struct {
	filepath        string
	linerange       string
	line            uint
	offset          uint
	start           uint
	end             uint
	displayLinenums bool
	unmatched       []string
}

func GetCliArgs() (string, dog.DogOptions, error) {
	cliArgs, err := getRawCliArgs()
	fmt.Printf("%+v\n", cliArgs)
	if err != nil {
		return "", dog.DogOptions{}, fmt.Errorf("error getting cli arguments\n%+v\n", err)
	}

	fpath, opts, err := processCliArgs(cliArgs)
	if err != nil {
		return "", dog.DogOptions{}, fmt.Errorf("error getting cli arguments\n%+v\n", err)
	}

	return fpath, opts, nil
}

func getRawCliArgs() (cliArgs, error) {
	// filepath '-f'
	fpath := flag.String("f", "", "path to target file")
	// range '-r'
	linerange := flag.String("r", "", "range of line numbers")
	// linenum, '-l'
	line := flag.Uint("l", 0, "line num of target")
	// offset, '-o'
	offset := flag.Uint("o", 0, "num of lines above and below")
	// start, '-s'
	startline := flag.Uint("s", 0, "line num of first line")
	// end, '-e'
	endline := flag.Uint("e", 0, "line num of last line")
	// display linenums '-n'
	linenums := flag.Bool("n", false, "display line nums")

	flag.Parse()

	return cliArgs{
		filepath:        *fpath,
		linerange:       *linerange,
		line:            *line,
		offset:          *offset,
		start:           *startline,
		end:             *endline,
		displayLinenums: *linenums,
		unmatched:       flag.Args(),
	}, nil
}

func processCliArgs(args cliArgs) (string, dog.DogOptions, error) {
	fpath := args.filepath
	// if filepath flag not found
	if fpath == "" {
		// check flag.Args
		if isEmpty(flag.Args()) { // empty, error
			return "", dog.DogOptions{}, fmt.Errorf("missing arg filepath")
		} else { // not empty, take first as filepath
			fpath = flag.Args()[0]
		}
	}

	// TODO: process other args

	return fpath, dog.DogOptions{}, nil
}

func isEmpty[T any](slice []T) bool {
	return len(slice) == 0
}
