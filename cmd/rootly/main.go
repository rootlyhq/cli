package main

import (
	"github.com/Matt-Gleich/lumber"
	"github.com/rootly-io/cli/pkg/commands"
)

func main() {
	lumber.ErrNilCheck = false
	commands.Execute()
}
