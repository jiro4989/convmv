package main

import (
	"fmt"
	"os"

	"github.com/docopt/docopt-go"
	"github.com/jiro4989/tkmvdbctl/commands/search"
	"github.com/jiro4989/tkmvdbctl/errorcode"
)

const (
	doc = `tkmvdbctl はツクールMVのデータベースファイルを操作するコマンドです。

Usage:
	tkmvdbctl <command> [<args>...]
	tkmvdbctl -h | --help
	tkmvdbctl -v | --version

Available commands:
	add    <target>
	remove <target>
	search <target> <word>

Options:
	-h --help                     Print this help.
	-v --version                  Print version.
	`
)

type (
	Config struct {
		Command string
		Args    []string
	}
)

func main() {
	os.Exit(int(Main(os.Args)))
}

func Main(argv []string) errorcode.ErrorCode {
	parser := &docopt.Parser{}
	args, _ := parser.ParseArgs(doc, argv[1:], Version)
	config := Config{}
	if err := args.Bind(&config); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return errorcode.FailedBinding
	}

	args2 := []string{argv[0], config.Command}
	args2 = append(args2, config.Args...)
	switch config.Command {
	case "add":
	case "remove":
	case "search":
		return search.CmdSearch(args2)
	}
	return errorcode.OK
}
