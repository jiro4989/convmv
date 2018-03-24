package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = "次郎(Jiro)"
	app.Version = "1.0.0"
	app.Author = "jiro4989"
	app.Email = ""
	app.Usage = "ツクールMVのデータをjsonからcsvといった表計算ソフトから編集しやすいデータ構造に相互変換します。"

	app.Flags = GlobalFlags
	app.Commands = Commands
	app.CommandNotFound = CommandNotFound

	app.Run(os.Args)
}
