package main

import (
	"fmt"
	"os"
	"os/user"
	"swagscript/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Whatup %s! Welcome to SwagScript! Go get yourself a bag\n",
		user.Username)
	fmt.Printf("Start typing in your commands!\n")
	repl.Start(os.Stdin, os.Stdout)
}
