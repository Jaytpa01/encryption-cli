package main

import (
	"log"

	"github.com/Jaytpa01/encryption-cli/cmd"
)

func main() {
	if err := cmd.RootCommand().Execute(); err != nil {
		log.Fatal(err)
		return
	}
}
