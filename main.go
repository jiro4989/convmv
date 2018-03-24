package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = Name
	app.Version = Version
	app.Author = "jiro4989"
	app.Email = ""
	app.Usage = "ツクールMVのデータをjsonからcsvといった表計算ソフトから編集しやすいデータ構造に相互変換します。"

	app.Flags = GlobalFlags
	app.Commands = Commands
	app.CommandNotFound = CommandNotFound

	app.Run(os.Args)
}
