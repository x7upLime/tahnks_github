package main

import (
	"fmt"
	"os"
	"os/user"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		fmt.Printf("There was an err: %q\n", err)
		os.Exit(-1)
	}

	fmt.Println("I am ", usr.Username)
	prev := os.Getenv("SUDO_USER")
	if prev != "" {
		fmt.Println("Previously I was ", prev)
	}
}
