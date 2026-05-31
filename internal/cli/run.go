package cli

import (
	"fmt"
	"os"

	"github.com/suhailkassar11/smalldocker/internal/runtime"
)


func Execute(args []string){

	if len(args)<2{
		fmt.Println("useage: smalldocker run <command>")
		os.Exit(1)
	}

	switch args[1]{
	case "run":
		runtime.Run(args[2:])
	case "child":
		runtime.Child(args[2:])

	default:
		println("runtime error")

	}
}