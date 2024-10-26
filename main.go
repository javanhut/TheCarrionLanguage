package main

import (
	"fmt"
	"os"
	"os/user"
	"thecarrionlang/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Carrion Programming Lanuage!\n", user.Username)
	fmt.Printf("Type any commands you like go with Odin\n")
	repl.Start(os.Stdin, os.Stdout)
}
