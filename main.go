package main

import (
	"gopkg.in/urfave/cli.v1"

	"os"
)

func main() {
	app := cli.NewApp()

	app.Run(os.Args)		
}