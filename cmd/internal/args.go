package internal

import (
	"flag"
	"fmt"

	"github.com/jaxfu/dog"
)

type CliArgs struct {
	Filepath        string
	Linerange       string
	Line            uint
	Offset          uint
	Start           uint
	End             uint
	DisplayLinenums bool
	Unmatched       []string
}

func GetCliArgs() (CliArgs, error) {
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

	return CliArgs{
		Filepath:        *fpath,
		Linerange:       *linerange,
		Line:            *line,
		Offset:          *offset,
		Start:           *startline,
		End:             *endline,
		DisplayLinenums: *linenums,
		Unmatched:       flag.Args(),
	}, nil
}

func ProcessCliArgs(args CliArgs) (string, dog.DogOptions, error) {
	fpath := ""
	// if filepath flag not found
	if args.Filepath == "" {
		// check flag.Args
		if isEmpty(flag.Args()) { // empty, error
			return "", dog.DogOptions{}, fmt.Errorf("missing arg filepath")
		} else { // not empty, take first as filepath
			fpath = flag.Args()[0]
		}
	} else {
	}

	return fpath, dog.DogOptions{}, nil
}

func isEmpty[T any](slice []T) bool {
	return len(slice) == 0
}
