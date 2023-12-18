package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/ricci2511/riccilang/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! Welcome to the riccilang programming language!\n", user.Username)
	fmt.Printf("Start playing around with it by typing in commands.\n")
	repl.Start(os.Stdin, os.Stdout)
}
