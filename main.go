package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/docopt/docopt-go"
	"github.com/jiro4989/tkmvdbctl/db"
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
	var m db.Map
	if err := json.Unmarshal(b, &m); err != nil {
		panic(err)
	}

	var result []TextEvent
	for _, evt := range m.Events {
		if evt == nil {
			continue
		}
		for _, page := range evt.Pages {
			var textFlag bool
			var textEvent TextEvent
			for _, l := range page.List {
				switch l.Code {
				case 101: // メタイベント
					if 0 < len(textEvent.Text) {
						result = append(result, textEvent)
						textEvent = TextEvent{}
					}
					textFlag = true
					if v, ok := l.Parameters[0].(string); ok {
						textEvent.Actor = v
					}
					if v, ok := l.Parameters[1].(int); ok {
						textEvent.ActorIndex = v
					}
					if v, ok := l.Parameters[2].(int); ok {
						textEvent.Background = v
					}
					if v, ok := l.Parameters[3].(int); ok {
						textEvent.Position = v
					}
				case 401: // テキストイベント
					if !textFlag {
						continue
					}
					p := l.Parameters[0]
					if t, ok := p.(string); ok {
						textEvent.Text = append(textEvent.Text, t)
					}
				default:
					if 0 < len(textEvent.Text) {
						result = append(result, textEvent)
						textEvent = TextEvent{}
					}
					textFlag = false
				}
			}
		}
	}
	for _, v := range result {
		for _, vv := range v.Text {
			fmt.Println(vv)
		}
	}
	return errorCodeOK
}

type TextEvent struct {
	Actor      string
	ActorIndex int
	Background int
	Position   int
	Text       []string
}
