package main

import (
  "strings"
  "net/url"
  "fmt"
)

func getid(u string) (string, error) {
  parse, err := url.Parse(u)
  if err != nil {
    return "", err
  }

  params, err := url.ParseQuery(parse.RawQuery)
  if err != nil {
    return "", err
  }

  id, ok := params["id"]
  if !ok || len(id) == 0 {
    return "", fmt.Errorf("IDを見つけられませんでした。")
  }

  return id[0], nil
}

/* 記事かの確認 */
func isarticle(u string) bool {
  chk := strings.Split(u, "=")
  return len(chk) > 2 && 
         (chk[0] == "/view_news.pl?id" || chk[0] == "/view_news.pl?from" || chk[0] == "/view_news.pl?media_id" || chk[0] == "/view_news.pl?stkt")
}

/* 部分圏かの確認 */
func issubcat(u string) bool {
  chk := strings.Split(u, "=")
  return len(chk) > 1 && 
         (chk[0] == "/list_news_category.pl?id" || chk[0] == "/list_news_category.pl?page" || chk[0] == "/list_news_category.pl?sort" || chk[0] == "/list_news_category.pl?type" || chk[0] == "/list_news_category.pl?sub_category_id") &&
         strings.Contains(u, "type=bn")
}

/* 部分かの確認 */
func iscategory(u string) bool {
  chk := strings.Split(u, "=")
  return len(chk) > 1 &&
         (chk[0] == "/list_news_category.pl?id" || chk[0] == "/list_news_category.pl?sub_category_id" || chk[0] == "/list_news_category?from") &&
         !strings.Contains(u, "type=bn")
}

/* 出版社かの確認 */
func ispublish(u string) bool {
  chk := strings.Split(u, "=")
  return len(chk) > 1 && (chk[0] == "/list_news_media.pl?id" || chk[0] == "/list_news_media.pl?page")
}

/* つぶやきかの確認 */
func istubayaki(u string) bool {
  chk := strings.Split(u, "=")
  return len(chk) > 1 &&
         (chk[0] == "/list_quote.pl?id" || chk[0] == "/list_quote.pl?type" || chk[0] == "/list_quote.pl?sort" || chk[0] == "/list_quote.pl?news_id") &&
         strings.Contains(u, "type=voice") &&
         (strings.Contains(u, "sort=post_time") || strings.Contains(u, "sort=feedback_count"))
}
