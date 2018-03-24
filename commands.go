package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/jiro4989/convmv/command"
)

var GlobalFlags = []cli.Flag{}

var Commands = []cli.Command{
	{
		Name:   "json2csv",
		Usage:  "ツクールMVのjsonデータをcsvデータに変換します。",
		Action: command.CmdJson2csv,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "csv2json",
		Usage:  "csvデータをツクールMVのjsonデータに変換します。",
		Action: command.CmdCsv2json,
		Flags:  []cli.Flag{},
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
