# dog

Go rewrite of the GNU coreutil [cat](https://www.gnu.org/software/coreutils/manual/html_node/cat-invocation.html)

## Arch

### Interface

- Filereader :: filepath string -> []string

- Processor :: \*buffer -> []string

### Types

- Line: repr of line of file
  - Address
  - Content

- Address: file name and linenum
  - Filename
  - Linenum
