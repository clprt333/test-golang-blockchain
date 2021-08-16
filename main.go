package main

import (
	"os"

	"github.com/clprt333/test-golang-blockchain/cli"
)

func main() {
	defer os.Exit(0)
	cmd := cli.CommandLine{}
	cmd.Run()
}
