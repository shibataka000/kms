package main

import (
	"log"

	"github.com/shibataka000/kms/cmd"
)

func main() {
	if err := cmd.NewRootCommand().Execute(); err != nil {
		log.Fatal(err)
	}
}
