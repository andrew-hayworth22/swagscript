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

	fmt.Printf("Hello %s! Welcome to the Monkey programming language!\n",
		user.Username)
	fmt.Printf("Start typing in your commands!\n")
	repl.Start(os.Stdin, os.Stdout)
}