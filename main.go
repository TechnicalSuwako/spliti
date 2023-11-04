package main

import (
  "fmt"
  "os"
  "strconv"
)

var version = "2.0.0"

func help() {
  fmt.Println("使い方：")
  fmt.Println("spliti -v               ：バージョンを表示")
  fmt.Println("spliti -s [ポート番号]　：ポート番号でウエブサーバーを実行（デフォルト＝9930）")
  fmt.Println("spliti -h               ：ヘルプを表示")
}

func main() {
  cnf, err := getconf()
  if err != nil {
    fmt.Println(err)
    return
  }

  args := os.Args

  if len(args) == 3 && args[1] == "-s" {
    if port, err := strconv.Atoi(args[2]); err != nil {
      fmt.Printf("%qは数字ではありません。\n", args[2])
      return
    } else {
      serv(cnf, port)
      return
    }
  } else if len(args) == 2 {
    if args[1] == "-v" {
      fmt.Println("spliti-" + version)
    } else if args[1] == "-s" {
      serv(cnf, 9930)
    }
  } else {
    help()
    return
  }
}
