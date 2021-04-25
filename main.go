//go:generate qtc -ext .html -dir views
//go:generate qtc -ext .sql -dir queries

package main

import (
	"os"

	"github.com/kyleu/admini/app/cmd"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		args = []string{"server"}
	}
	cmd.Run(args[0], args[1:])
}
