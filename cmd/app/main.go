package main

import (
	"os"

	"kvstore/internal/repl"
	"kvstore/internal/store"
)

func main() {
	s := store.New()
	repl.Run(os.Stdin, os.Stdout, s)
}

