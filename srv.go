package main

import (
  "text/template"
  "fmt"
  "net/http"
  "net/url"
  "strings"
  "log"
  "os"
  "path/filepath"
)

type Page struct {
  Tit, Err, Bdy, Dec, Img, Url, Dom, Ver, Ves string
}

func extractGurl(r *http.Request) (string, error) {
  rq := r.URL.RawQuery
  q, err := url.QueryUnescape(rq)
  if err != nil {
    return "URLを受取に失敗", err
  }

  gurl := strings.Replace(q, "/?url=", "", -1)

  return gurl, nil
}

func serv (cnf Config, port int) {
  dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
  if err != nil {
    log.Fatal(err)
  }
  err = os.Chdir(dir)
  if err != nil {
    log.Fatal(err)
  }

  http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(cnf.webpath + "/static"))))
  ftmpl := []string{cnf.webpath + "/view/index.html", cnf.webpath + "/view/header.html", cnf.webpath + "/view/footer.html"}

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    // 1.xで、URLは「/?url=」が付いたけど、2.0.0からは不要になった
    // だから、下記の部分は古いURLの為だ
    urls, ok := r.URL.Query()["url"]
    if ok && len(urls[0]) > 0 {
      http.Redirect(w, r, "/" + urls[0], http.StatusMovedPermanently)
      return
    }

    data := &Page{Ver: version, Ves: strings.ReplaceAll(version, ".", "")}
    uri := r.URL.Path
    gurl, err := extractGurl(r)
    if err != nil {
      data.Tit = "エラー"
      data.Err = err.Error()
      ftmpl[0] = cnf.webpath + "/view/404.html"
    }

    if uri == "/" {
      ftmpl[0] = cnf.webpath + "/view/index.html"
    } else {
      furl := uri + "?" + gurl
      page := get(furl, cnf)
      data.Tit = page["title"]
      if page["err"] != "" {
        data.Err = page["err"]
        ftmpl[0] = cnf.webpath + "/view/404.html"
      } else {
        data.Bdy = page["content"]
        data.Img = "/static/logo.jpg"
        if isarticle(furl) {
          data.Dec = page["desc"]
          data.Img = page["img"]
          data.Url = cnf.domain + furl
        }
        ftmpl[0] = cnf.webpath + "/view/news.html"
      }
    }

    tmpl := template.Must(template.ParseFiles(ftmpl[0], ftmpl[1], ftmpl[2]))
    tmpl.Execute(w, data)
    data = nil
  })

  fmt.Println(fmt.Sprint("http://" + cnf.ip + ":", port, " でサーバーを実行中。終了するには、CTRL+Cを押して下さい。"))
  http.ListenAndServe(fmt.Sprint(cnf.ip + ":", port), nil)
}
