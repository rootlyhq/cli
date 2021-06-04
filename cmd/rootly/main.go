package main

import (
	"github.com/gleich/lumber"
	"github.com/rootly-io/cli/pkg/commands"
)

func main() {
	lumber.ShowStack = false
	lumber.ErrNilCheck = false
	commands.Execute()
}
