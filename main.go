package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/docopt/docopt-go"
)

const (
	doc = `tkmvdbctl はツクールMVのデータベースファイルを操作するコマンドです。

Usage:
	tkmvdbctl <command> <target> [options]
	tkmvdbctl -h | --help
	tkmvdbctl -v | --version

Available commands:
	add    <target>
	remove <target>
	search <target>

Available targets:
	text
	skill
	weapon
	armor

Options:
	-h --help                     Print this help.
	-v --version                  Print version.

Message options:
	-a --actor=<ACTOR>              Actor image.
	-m --message=<MESSAGE>          Message.
	-b --background=<BACKGROUND>    Window background color.
	-p --position=<POSITION>        Window position.
	-f --file=<FILE>                Event.json path.
	`
)

const (
	errorCodeOK ErrorCode = iota
	errorCodeFailedBinding
	errorCodeCouldNotReadFile
	errorCodeIllegalMonitorConfig
	errorCodeCouldNotReadDir
)

type (
	Config struct {
		Command    string
		Target     string
		Actor      string
		Message    string
		Background string
		Position   string
		File       string
	}
	ErrorCode int
)

func main() {
	os.Exit(int(Main(os.Args)))
}

func Main(argv []string) ErrorCode {
	parser := &docopt.Parser{}
	args, _ := parser.ParseArgs(doc, argv[1:], Version)
	config := Config{}
	if err := args.Bind(&config); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return errorCodeFailedBinding
	}

	switch config.Command {
	case "add":
	case "remove":
	case "search":
		return CmdSearch(config)
	}
	return errorCodeOK
}

func CmdSearch(c Config) ErrorCode {
	switch c.Target {
	case "text":
		return SearchText(c)
	case "skill":
	case "weapon":
	case "armor":
	}
	return errorCodeOK
}

func SearchText(c Config) ErrorCode {
	jsonFile := "/home/jiro4989/Documents/Games/Project1/data/Map001.json"
	b, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		panic(err)
	}
	fmt.Println(b)
	//json.Unmarshal(b)
	return errorCodeOK
}
