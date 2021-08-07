package main

import (
	"github.com/gleich/lumber"
	"github.com/rootlyhq/cli/pkg/commands"
)

func main() {
	lumber.ShowStack = false
	lumber.ErrNilCheck = false
	commands.Execute()
}
