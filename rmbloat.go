package main

import (
  "regexp"
  "strings"
)

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
