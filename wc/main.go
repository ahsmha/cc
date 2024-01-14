package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// todo - reformatting and proper functions and chars logic
	// bytes.Fields
	// process cmd args
	args := os.Args
	// help str print
	helpStr := `usage: wc [options] filepath

yet another wc.

options:
    -c | --chars                create a new markdown post
    -l | --lines 				create a new markdown post
    -w | --words				create a new markdown post
`
	var filename string
	if len(args) > 1 {
		filename = args[2]
	} else {
		fmt.Println(helpStr)
		os.Exit(1)
	}

	switch args[1] {
	case "-c", "--chars":
		bf, err := os.ReadFile(filename)
		if err != nil {
			os.Exit(1)
		}

		fmt.Println(len(bf))
		fmt.Println(len(strings.Split(string(bf), "\n")))
		fmt.Println(len(strings.FieldsFunc(string(bf), func(r rune) bool {
			return r == ' ' || r == '\t' || r == '\n' || r == '\f' || r == '\v' || r == '\r'
		})))
	case "-l", "--lines":
	case "-w", "--words":
	default:
		fmt.Println("option not found. please try again.")
	}
	return
}
