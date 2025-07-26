package shared

type Line struct {
	Address Address
	Content string
}

type Address struct {
	Filepath string
	Linenum  uint
}
