package cmd

import (
	"log"
	"os"
)

// Execute starts the program.
func Execute() {
	cmdBuilder := commandBuilder{}
	rootCommand := cmdBuilder.build()
	if err := rootCommand.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
