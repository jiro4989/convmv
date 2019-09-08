package search

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/docopt/docopt-go"
	"github.com/jiro4989/tkmvdbctl/db"
	"github.com/jiro4989/tkmvdbctl/errorcode"
)

const (
	doc = `Usage:
	tkmvdbctl search [options] <target> <word>
	tkmvdbctl search -h | --help

Available targets:
	text
	skill
	weapon
	armor

Options:
	-h --help                     Print this help.
	-F --format=<FORMAT>          Print format. (plain | json) [default: plain]
	-a --actor=<ACTOR>              Actor image.
	-m --message=<MESSAGE>          Message.
	-b --background=<BACKGROUND>    Window background color.
	-p --position=<POSITION>        Window position.
	-f --file=<FILE>                Event.json path.
	`
)

type (
	Config struct {
		Command    bool `docopt:"search"`
		Format     string
		Target     string
		Word       string
		Actor      string
		Message    string
		Background string
		Position   string
		File       string
	}
	TextEvent struct {
		Actor      string
		ActorIndex int
		Background int
		Position   int
		Text       []string
	}
)

func CmdSearch(argv []string) errorcode.ErrorCode {
	parser := &docopt.Parser{}
	args, _ := parser.ParseArgs(doc, argv[1:], "")
	c := Config{}
	if err := args.Bind(&c); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return errorcode.FailedBinding
	}

	switch c.Target {
	case "text":
		return SearchText(c)
	case "skill":
	case "weapon":
	case "armor":
	}
	return errorcode.OK
}

func SearchText(c Config) errorcode.ErrorCode {
	jsonFile := "/home/jiro4989/Documents/Games/Project1/data/Map001.json"
	b, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		panic(err)
	}
	var m db.Map
	if err := json.Unmarshal(b, &m); err != nil {
		panic(err)
	}

	var textEvents []TextEvent
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
						textEvents = append(textEvents, textEvent)
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
						textEvents = append(textEvents, textEvent)
						textEvent = TextEvent{}
					}
					textFlag = false
				}
			}
		}
	}

	// 検索ワードとマッチするものに絞り込み
	var matched []TextEvent
	for _, v := range textEvents {
		for _, t := range v.Text {
			if strings.Contains(t, c.Word) {
				matched = append(matched, v)
				break
			}
		}
	}

	for _, v := range matched {
		fmt.Println(v)
	}

	return errorcode.OK
}
