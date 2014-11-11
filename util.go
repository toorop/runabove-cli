package main

import (
	"fmt"
	"os"
	//"runtime/debug"
)

// dieError handle fatal error
func dieError(v ...interface{}) {
	fmt.Print("ERROR: ")
	for _, vv := range v {
		fmt.Printf("%v", vv)
	}
	fmt.Println("")
	//debug.PrintStack()
	os.Exit(1)
}

func dieOk(r ...string) {
	if len(r) != 0 {
		for _, line := range r {
			fmt.Println(line)
		}
	}
	os.Exit(0)
}

// Exit & and display error on bad arguments
func dieBadArgs(msg ...string) {
	errMsg := "Bad arg(s). Run ./ra command [subCommand...] --help for help."
	if len(msg) > 0 {
		errMsg = msg[0]
	}
	dieError(errMsg)
}

// Exit if args are missing
func dieIfArgsMiss(nbArgs, requiered int) {
	if nbArgs < requiered {
		dieBadArgs()
	}
}

func dieOnError(err error) {
	if err != nil {
		dieError(err)
	}
}
