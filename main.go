package main

import (
	"CardHero/ch"
	"fmt"
	"log"
	"os"
)

func main() {
	rohit, err := ch.NewUser("Rohit", "Awate", "awate.r@northeastern.edu", "hello123")
	if err != nil {
		log.Println("Invalid email!")
		os.Exit(1)
	}

	fmt.Println(ch.NewCard(*rohit, "hello, world!"))
}
