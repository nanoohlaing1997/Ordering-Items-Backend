package main

import (
	"log"
)

func main() {
	// Register Cli
	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
