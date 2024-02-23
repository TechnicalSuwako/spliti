package main

import (
  "fmt"
  "io"
  "net/http"
  "regexp"
  "strings"
)

/* ページのタイトル */
func gettitle(str string) string {
  re := regexp.MustCompile("<title>(.*)</title>")
  matches := re.FindStringSubmatch(str)
  if len(matches) > 1 {
    return matches[1]
  }
  return ""
}

func getimg(str string, cnf Config) string {
  re := regexp.MustCompile(`<img class="NEWS_tempPhoto__picture" src="(.*)" alt="">`)
  matches := re.FindStringSubmatch(str)
  if len(matches) > 1 {
    return strings.Replace(matches[1], "https://", cnf.imgproxy+"/", -1)
  }
  return ""
}

func getdesc(str string) string {
  re := regexp.MustCompile(`<div class="newsArticle">(.*?)</div>`)
  res := re.ReplaceAllString(str, "")
  return strip_tags(res)
}

/* 記事の受取 */
func get(url string, cnf Config) map[string]string {
  // デフォルト＝エラー
  res := make(map[string]string)
  res["title"] = "見つけられない"
  res["content"] = `
  <div class="newsArticle"><div class="articleHeading02">
    <div class="headingArea">
      <h1>見つけられなかった</h1>
    </div>
  </div>
  <div class="contents clearfix">
    <div class="article decoratable">
      <p>ごめんね！</p>
    </div>
  </div>
  `
  res["img"] = ""
  res["desc"] = ""
  res["err"] = ""

  resp, err := http.Get("https://news.mixi.jp" + url)
  if err != nil {
    res["err"] = "URLエラー"
    fmt.Println(res["err"] + ": " + err.Error())
    return res
  }
  defer resp.Body.Close()

  if resp.StatusCode == http.StatusOK {
    bytebody, err := io.ReadAll(resp.Body)
    if err != nil {
      res["err"] = "内容はバイトコードとして読み込みに失敗。"
      fmt.Println(res["err"])
      return res
    }

    body, err := EUCJPToUTF8(bytebody)
    if err != nil {
      res["err"] = err.Error()
      fmt.Println(res["err"])
      return res
    }

    id, _ := getid(url)

    res["title"] = gettitle(body)
    if isarticle(url) {
      if !strings.Contains(body, "newsArticle") {
        res["content"] = rmebloat(body, cnf)
      } else {
        res["img"] = getimg(body, cnf)
        res["content"] = rmbloat(id, body, cnf)
      }
    } else if ispublish(url) {
      res["content"] = rmpbloat(body, cnf)
    } else if issubcat(url) {
      if !strings.Contains(body, "subCategoryNavi") {
        res["content"] = rmebloat(body, cnf)
      } else {
        res["content"] = rmsbloat(body, cnf)
      }
    } else if istubayaki(url) {
      if !strings.Contains(body, "quoteList") {
        res["content"] = rmebloat(body, cnf)
      } else {
        res["content"] = rmqbloat(body, cnf)
      }
    } else {
      if !strings.Contains(body, "注目のニュース") {
        res["content"] = rmebloat(body, cnf)
      } else {
        res["content"] = rmcbloat(body, cnf)
      }
    }
    res["desc"] = getdesc(res["content"])
  }

  return res
}
