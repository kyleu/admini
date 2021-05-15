package main

import (
	"github.com/kyleu/admini/app/cmd"
)

func main() {
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
