package main

import (
	"os"

	"github.com/seungyongshim/countbeat/cmd"

	_ "github.com/seungyongshim/countbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
