package main

import (
	"fmt"
	"os"
	"os/user"

	"interpreter/repl"
)

func main() {
	user, err := user.Current()

	if err != nil {
		fmt.Println("Error finding current user")
		return
	}

	fmt.Printf("Hi %s, welcome to the REPL\n", user.Username)
	fmt.Println("Type any command to get started.")

	repl.Start(os.Stdin, os.Stdout)
}
