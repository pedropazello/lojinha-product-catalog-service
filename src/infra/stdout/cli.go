package stdout

import (
	"fmt"
	"os"
)

type CLI struct {
}

func (c *CLI) GetStdInput() []string {
	return os.Args[1:]
}

func (c *CLI) PutStdOutput(msg string) {
	fmt.Println(msg)
}
