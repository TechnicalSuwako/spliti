package main

import (
  "fmt"
  "encoding/json"
  "io/ioutil"
  "runtime"
  "os"
  "errors"
  "strings"
)

type Config struct {
  configpath, webpath, domain, imgproxy, ip string
}

func getconf () (Config, error) {
  var cnf Config

  prefix := "/usr"
  if runtime.GOOS == "freebsd" || runtime.GOOS == "openbsd" {
    prefix += "/local"
  } else if runtime.GOOS == "netbsd" {
    prefix += "/pkg"
  }

  cnf.configpath = "/etc/spliti/config.json"
  if runtime.GOOS == "freebsd" || runtime.GOOS == "netbsd" {
    cnf.configpath = prefix + cnf.configpath
  }

  data, err := ioutil.ReadFile(cnf.configpath)
  if err != nil {
    fmt.Println("config.jsonを開けられません: ", err)
    return cnf, errors.New("コンフィグファイルは " + cnf.configpath + " に創作して下さい。")
  }

  var payload map[string]interface{}
  json.Unmarshal(data, &payload)
  if payload["webpath"] == nil {
    payload["webpath"] = "/var/www/htdocs/spliti"
  }
  if payload["domain"] == nil {
    return cnf, errors.New("「domain」の値が設置していません。")
  }
  if payload["imgproxy"] == nil {
    payload["imgproxy"] = "https://imgproxy.076.moe"
  }
  if payload["ip"] == nil {
    payload["ip"] = "0.0.0.0"
  }
  if _, err := os.Stat(payload["webpath"].(string)); err != nil {
    fmt.Printf("%v\n", err)
    return cnf, errors.New("mkdirコマンドを使って、 " + payload["webpath"].(string))
  }
  if !strings.HasPrefix(payload["domain"].(string), "http://") && !strings.HasPrefix(payload["domain"].(string), "https://") {
    return cnf, errors.New("URLは「http://」又は「https://」で始める様にして下さい。")
  }
  cnf.webpath = payload["webpath"].(string)
  cnf.domain = payload["domain"].(string)
  cnf.imgproxy = payload["imgproxy"].(string)
  cnf.ip = payload["ip"].(string)
  payload = nil

  return cnf, nil
}
