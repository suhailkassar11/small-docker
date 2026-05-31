package app

import (
	"os"

	"github.com/suhailkassar11/smalldocker/internal/cli"
)

func Run(){
	cli.Execute(os.Args)
}