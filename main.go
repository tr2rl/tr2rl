package main

//go:generate go-winres make

import "github.com/cytificlabs/tr2rl/cmd"

func main() {
	cmd.Execute()
}
