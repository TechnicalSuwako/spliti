package main

import (
  "fmt"
  "io/ioutil"
  "net/http"
  "regexp"
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

/* 記事かの確認 */
func isarticle(url string) bool {
  chk := strings.Split(url, "=")
  return len(chk) > 2 && 
         (chk[0] == "/view_news.pl?id" || chk[0] == "/view_news.pl?from" || chk[0] == "/view_news.pl?media_id" || chk[0] == "/view_news.pl?stkt")
}

/* 部分圏かの確認 */
func issubcat(url string) bool {
  chk := strings.Split(url, "=")
  return len(chk) > 1 && 
         (chk[0] == "/list_news_category.pl?id" || chk[0] == "/list_news_category.pl?page" || chk[0] == "/list_news_category.pl?sort" || chk[0] == "/list_news_category.pl?type" || chk[0] == "/list_news_category.pl?sub_category_id") &&
         strings.Contains(url, "type=bn")
}

/* 部分かの確認 */
func iscategory(url string) bool {
  chk := strings.Split(url, "=")
  return len(chk) > 1 &&
         (chk[0] == "/list_news_category.pl?id" || chk[0] == "/list_news_category.pl?sub_category_id" || chk[0] == "/list_news_category?from") &&
         !strings.Contains(url, "type=bn")
}

/* 出版社かの確認 */
func ispublish(url string) bool {
  chk := strings.Split(url, "=")
  return len(chk) > 1 && (chk[0] == "/list_news_media.pl?id" || chk[0] == "/list_news_media.pl?page")
}

/* カテゴリーだけが残るまで消す */
func rmcbloat(body string, cnf Config) string {
	var re *regexp.Regexp

	rep := []struct {
		pat  string
		repl string
	}{
		{`(?s)<!DOCTYPE html>.*?<!--注目のニュース-->`, ""},
    {`(?s)<!--/newsCategoryList-->.*?</html>`, ""},
    {`(?s)<div class="COMMONDOC_header2017_mainNavHeader.*?</div>`, ""},
    {`(?s)<div class="COMMONDOC_header2017_mainNav.*?</div>`, ""},
    {`(?s)<div class="COMMONDOC_header2017_toggleNav.*?</div>`, ""},
    {`(?s)<div class="COMMONDOC_header2017_localNavArea.*?</div>`, ""},
    {`(?s)<div class="COMMONDOC_header2017_globalNav__account.*?</div>`, ""},
    {`(?s)<div class="COMMONDOC_header2017_globalNav__logo.*?</div>`, ""},
    {`(?s)<div class="COMMONDOC_header2017_globalNav__toggleNav.*?</div>`, ""},
    {`(?s)<div class="COMMONDOC_header2017_adBanner.*?</div>`, ""},
    {`(?s)<div class="COMMONDOC_header2017_globalNav.*?</div>`, ""},
    {`(?s)<div class="COMMONDOC_header2017_globalNavArea.*?</div>`, ""},
    {`(?s)<div class="COMMONDOC_header2017_headerArea.*?</div>`, ""},
    {`(?s)<div id="page" class="FRAME2016_page">.*?</div>`, ""},
    {`(?s)<div id="subCategoryNavi.*?</div>`, ""},
    {`(?s)<div role="navigation".*?</div>`, ""},
    {`(?s)<div id="div-gpt-ad-.*?</div>`, ""},
    {`(?s)<h3.*?</h3>`, ""},
    {`(?s)<script.*?</script>`, ""},
    {`(?s)<ul class="entryList0.*?</ul>`, ""},
    {`(?s)<div class="adMain.*?</div>`, ""},
    {`(?s)<div class="gAdComponent.*?</div>`, ""},
    {`(?s)<div class="adsense0.*?</div>`, ""},
    {`(?s)<div class="adsense.*?</div>`, ""},
    {`(?s)<div class="pageList02.*?</div>`, ""},
    {`(?s)<span class="reactionCountBalloon.*?</span>`, ""},
		{`https://news-image.mixi.net`, cnf.imgproxy + `/news-image.mixi.net`},
		{`https://img.mixi.net`, cnf.imgproxy + `/img.mixi.net`},
		{`https://news.mixi.jp/`, cnf.domain + `/`},
    {`・ `, ""},
    {`\[`, ""},
    {`\]`, ""},
	}

	for _, r := range rep {
		re = regexp.MustCompile(r.pat)
		body = re.ReplaceAllString(body, r.repl)
	}

  body = strings.TrimSpace("<div class=\"subCategoryNavi\" class=\"LEGACY_UI2016_subCategoryNavi\">\n" + strings.TrimSpace(body)) + "\n    </div>\n"
	return "<div class=\"newsArticle\">" + body + "</div>"
}

/* エラーだけが残るまで消す */
func rmebloat(body string, cnf Config) string {
	var re *regexp.Regexp

	rep := []struct {
		pat  string
		repl string
	}{
		{`(?s)<!DOCTYPE html>.*?<p class="messageAlert">`, ""},
		{`(?s)</p>.*?</html>`, ""},
	}

	for _, r := range rep {
		re = regexp.MustCompile(r.pat)
		body = re.ReplaceAllString(body, r.repl)
	}

	body = strings.TrimSpace("<div class=\"newsArticle\">\n" + strings.TrimSpace(body)) + "\n    </div>\n"
	return body
}

/* 部分圏だけが残るまで消す */
func rmsbloat(body string, cnf Config) string {
  var re *regexp.Regexp

  rep := []struct {
    pat  string
    repl string
  }{
    {`(?s)<!DOCTYPE html>.*?<!-- InstanceBeginEditable name="bodyMain" -->`, ""},
    {`(?s)<div class="adsenseBannerArea">.*?</html>`, ""},
    //{`(?s)<div class="pageList02.*?</div>`, ""},
		{`https://news-image.mixi.net`, cnf.imgproxy + `/news-image.mixi.net`},
		{`https://img.mixi.net`, cnf.imgproxy + `/img.mixi.net`},
		{`https://news.mixi.jp/`, cnf.domain + `/`},
    {`・ `, ""},
    {`\[`, ""},
    {`\]`, ""},
  }

  for _, r := range rep {
    re = regexp.MustCompile(r.pat)
    body = re.ReplaceAllString(body, r.repl)
  }

	body = strings.TrimSpace("<div class=\"newsArticle\">\n" + strings.TrimSpace(body)) + "\n    </div>\n"
	return body
}

/* 出版社だけが残るまで消す */
func rmpbloat(body string, cnf Config) string {
  var re *regexp.Regexp

  rep := []struct {
    pat  string
    repl string
  }{
    {`(?s)<!DOCTYPE html>.*?<!-- InstanceBeginEditable name="bodyMain" -->`, ""},
    {`(?s)<!-- InstanceEndEditable -->.*?</html>`, ""},
    {`(?s)<div class="pageList02.*?</div>`, ""},
		{`https://news-image.mixi.net`, cnf.imgproxy + `/news-image.mixi.net`},
		{`https://img.mixi.net`, cnf.imgproxy + `/img.mixi.net`},
		{`https://news.mixi.jp/`, cnf.domain + `/`},
    {`・ `, ""},
    {`\[`, ""},
    {`\]`, ""},
  }

  for _, r := range rep {
    re = regexp.MustCompile(r.pat)
    body = re.ReplaceAllString(body, r.repl)
  }

	body = strings.TrimSpace("<div class=\"newsArticle\">\n" + strings.TrimSpace(body)) + "\n    </div>\n"
	return body
}

/* 記事だけが残るまで消す */
func rmbloat(body string, cnf Config) string {
	var re *regexp.Regexp

	rep := []struct {
		pat  string
		repl string
	}{
		{`(?s)<!DOCTYPE html>.*?<div class="newsArticle">`, ""},
		{`(?s)<!--/newsArticle-->.*?</html>`, ""},
		{`(?s)<p class="reactions">.*?</p>`, ""},
		{`(?s)<ul class="diaryUtility\d*">.*?</ul>`, ""},
		{`(?s)<table>.*?</table>`, ""},
		{`(?s)<div class="adsense0.*?</div>`, ""},
		{`(?s)<div class="adsense.*?</div>`, ""},
		{`www\.?youtube\.com`, "youtube.owacon.moe"},
		{`(?s)<div class="subInfo">.*?</div>`, ""},
		{`(?s)<div class="additional\d*.*?</div>`, ""},
		{`(?s)(^[\r\n]*|[\r\n]+)[\s\t]*[\r\n]+`, "\n"},
		{`<!--article_image-->`, ""},
		{`<!--/article_image-->`, ""},
		{`(?s)<!--.*?-->`, ""},
		{`<!--`, ""},
		{`(?s)<img src="https://(.*?)"`, `<img src="` + cnf.imgproxy + `/$1"`},
		{`https://news-image.mixi.net`, cnf.imgproxy + `/news-image.mixi.net`},
		{`https://news.mixi.jp/`, cnf.domain + `/`},
	}

	for _, r := range rep {
		re = regexp.MustCompile(r.pat)
		body = re.ReplaceAllString(body, r.repl)
	}

	body = strings.TrimSpace("<div class=\"newsArticle\">\n" + strings.TrimSpace(body)) + "\n    </div>\n"
	return body
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
    bytebody, err := ioutil.ReadAll(resp.Body)
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

    res["title"] = gettitle(body)
    if isarticle(url) {
      if !strings.Contains(body, "newsArticle") {
        res["content"] = rmebloat(body, cnf)
      } else {
        res["img"] = getimg(body, cnf)
        res["content"] = rmbloat(body, cnf)
      }
    } else if ispublish(url) {
      res["content"] = rmpbloat(body, cnf)
    } else if issubcat(url) {
      if !strings.Contains(body, "subCategoryNavi") {
        res["content"] = rmebloat(body, cnf)
      } else {
        res["content"] = rmsbloat(body, cnf)
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
