package main

import (
  "io/ioutil"
  "strings"
  "bytes"

  "golang.org/x/net/html"
  "golang.org/x/text/encoding/japanese"
  "golang.org/x/text/transform"
)

/* PHPであるstrip_tagsはGo言語で存在しないから、自分で作る */
func strip_tags(data string) string {
  doc, err := html.Parse(strings.NewReader(data))
  if err != nil {
    panic("HTMLをパーシングに失敗。")
  }

  var buf bytes.Buffer
  var f func(*html.Node)
  f = func(n *html.Node) {
    if n.Type == html.TextNode {
      buf.WriteString(n.Data)
    }
    for c := n.FirstChild; c != nil; c = c.NextSibling {
      f(c)
    }
  }
  f(doc)

  return buf.String()
}

func EUCJPToUTF8(input []byte) (string, error) {
  transformer := japanese.EUCJP.NewDecoder()
  reader := transform.NewReader(bytes.NewReader(input), transformer)
  result, err := ioutil.ReadAll(reader)
  if err != nil {
    return "エンコーディングに失敗", err
  }

  return string(result), nil
}
