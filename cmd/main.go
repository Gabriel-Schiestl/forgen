package main

import "github.com/Gabriel-Schiestl/forgen/cmd/command"

func main() {
	if err := command.RootCmd.Execute(); err != nil {
		panic(err)
	}
}
