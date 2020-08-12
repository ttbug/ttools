package main

import (
	"log"

	"github.com/ttbug/ttools/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}

}
