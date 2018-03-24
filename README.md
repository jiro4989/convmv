# convmv (convert-mv)

***開発途中***

ツクールMVのデータをjsonからCSVといったExcelやOpenOfficeから操作しやすいデータ構
造に変換します。また、変換したCSVを元のjsonに変換する機能を有します。

## 目的

ツクールMVのUIは大量の似たようなデータを作成するのに適しておりません。ツクールMV
のエディタの問題点は下記のとおりです。

- 一つのデータを画面内に一つしか表示できないため、隣接するデータとの比較ができな
  い。
- データのフィルタができない。
- 同じスキルのランク違いといったスキルを大量に作るのに適さない。
- データの検索ができない

上記の問題点は既存の表計算ソフトの Excel または LibreOffice といったソフトを使用
することで解消可能と考えております。

しかし、どちらのソフトもJson形式のインポートをサポートしておらず、またインポート
可能でもゲーム内のデータと表示上のデータのひも付けが必要であり、それらの機能は当
然どちらのソフトもサポートしておりません。

本ツールは、上記の表計算ソフトから操作しやすいデータ構造に変換する部分と、ゲーム
内コードと表示上のデータのひも付けの部分をサポートします。

## 動作環境

下記の環境での動作をサポートします。

| OS      | アーキテクチャ |
|:--------|:---------------|
| Windows | 32bit          |
| Windows | 64bit          |
| Mac     | 32bit          |
| Mac     | 64bit          |
| Linux   | 32bit          |
| Linux   | 64bit          |

## 使い方

本リポジトリの成果物の下記２つをツクールのフォルダに配置します。

- convmv-json2csv.exe
- convmv-csv2json.exe

ツクールのフォルダに配置した際のフォルダ構造は下記のようになります。

TODO

    tkool_project
    ├ audio/
    ├ save/
    ├ data/
    ├ fonts/
    ├ movies/
    ├ js/
    ├ icon/
    ├ img/
    ├ index.html
    └ Game.rpgproject

convmv-json2csv.exeをダブルクリックします。

dataフォルダ内にconvmv-skill.csvが生成されることを確認します。

Excelなどの表計算ソフトで生成されたCSVを編集します。

convmv-csv2jsonをダブルクリックします。

ツクールMVを起動します。

編集したデータを確認し、更新されていることを確認します。

## インストール

To install, use `go get`:

```bash
$ go get -d github.com/jiro4989/convmv
```

## 作者

[次郎(Jiro)](https://github.com/jiro4989)
