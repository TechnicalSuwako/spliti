package main

import (
  "strings"
)

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
