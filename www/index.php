<?php
  include("../config.php");

  /* ページのタイトル */
  function gettitle (string $str): string {
    preg_match("/<title>(.*)<\/title>/", $str, $matches);
    return $matches[1];
  }

  function getimg (string $str): string {
    preg_match('/<img class="NEWS_tempPhoto__picture" src="(.*)" alt="">/', $str, $matches);
    $res = str_replace("https://", IMGPROXY."/", $matches[1]);
    unset($matches);
    return $res;
  }

  function getdesc (string $str): string {
    $res = preg_replace('/<div class="newsArticle">(.*?)<\/div>/s', "", $str);
    return strip_tags($res);
  }

  /* カテゴリーかの確認 */
  function iscategory (string $url): bool {
    $chk = explode("=", $url);
    if (isset($chk[0])) $chk = explode("?", $chk[0]);
    return isset($chk[0]) && $chk[0] == "list_news_category.pl";
  }

  /* 記事かの確認 */
  function isarticle (string $url): bool {
    $chk = explode("=", $url);
    return isset($chk[0]) && $chk[0] == "view_news.pl?id" && isset($chk[1]) && isset($chk[2]);
  }

  /* カテゴリーだけが残るまで消す */
  function rmcbloat (string $body): string {
    //$res = preg_replace('/<!DOCTYPE html>(.*?)<div id="subCategoryNavi" class="LEGACY_UI2016_subCategoryNavi">/s', "", $body);
    $res = preg_replace('/<!DOCTYPE html>(.*?)<!--注目のニュース-->/s', "", $body);
    $res = preg_replace('/<!--\/newsCategoryList-->(.*?)<\/html>/s', "", $res);
    $res = preg_replace('/<div class="COMMONDOC_header2017_mainNavHeader(.*?)<\/div>/s', "", $res);
    $res = preg_replace('/<div class="COMMONDOC_header2017_mainNav(.*?)<\/div>/s', "", $res);
    $res = preg_replace('/<div class="COMMONDOC_header2017_toggleNav(.*?)<\/div>/s', "", $res);
    $res = preg_replace('/<div class="COMMONDOC_header2017_localNavArea(.*?)<\/div>/s', "", $res);
    $res = preg_replace('/<div class="COMMONDOC_header2017_globalNav__account(.*?)<\/div>/s', "", $res);
    $res = preg_replace('/<div class="COMMONDOC_header2017_globalNav__logo(.*?)<\/div>/s', "", $res);
    $res = preg_replace('/<div class="COMMONDOC_header2017_globalNav__toggleNav(.*?)<\/div>/s', "", $res);
    $res = preg_replace('/<div class="COMMONDOC_header2017_adBanner(.*?)<\/div>/s', "", $res);
    $res = preg_replace('/<div class="COMMONDOC_header2017_globalNav(.*?)<\/div>/s', "", $res);
    $res = preg_replace('/<div class="COMMONDOC_header2017_globalNavArea(.*?)<\/div>/s', "", $res);
    $res = preg_replace('/<div class="COMMONDOC_header2017_headerArea(.*?)<\/div>/s', "", $res);
    $res = preg_replace('/<div id="page" class="FRAME2016_page">(.*?)<\/div>/s', "", $res);
    $res = preg_replace('/<div id="subCategoryNavi(.*?)<\/div>/s', "", $res);
    $res = preg_replace('/<div role="navigation"(.*?)<\/div>/s', "", $res);
    $res = preg_replace('/<div id="div-gpt-ad-(.*?)<\/div>/s', "", $res);
    $res = preg_replace('/<h3(.*?)<\/h3>/s', "", $res);
    $res = preg_replace('/<script(.*?)<\/script>/s', "", $res);
    $res = preg_replace('/<ul class="entryList0(.*?)<\/ul>/s', "", $res);
    $res = preg_replace('/<div class="adMain(.*?)<\/div>/s', "", $res);
    $res = preg_replace('/<div class="gAdComponent(.*?)<\/div>/s', "", $res);
    $res = preg_replace('/<div class="adsense0(.*?)<\/div>/s', "", $res);
    $res = preg_replace('/<div class="adsense(.*?)<\/div>/s', "", $res);
    $res = preg_replace('/<div class="pageList02(.*?)<\/div>/s', "", $res);
    $res = preg_replace('/<span class="reactionCountBalloon(.*?)<\/span>/s', "", $res);
    $res = str_replace("https://news-image.mixi.net", IMGPROXY."/news-image.mixi.net", $res);
    $res = str_replace("https://img.mixi.net", IMGPROXY."/img.mixi.net", $res);
    $res = str_replace("https://news.mixi.jp/", DOMAIN."/?url=", $res);
    $res = str_replace("・ ", "", $res);
    $res = str_replace("[", "", $res);
    $res = str_replace("]", "", $res);
    //$res = preg_replace("/sort=(.*?)&page=(.*?)&id=(.*?)/", "id=$0&page=$1&sort=$2", $res);
    $res = trim("<div id=\"subCategoryNavi\" class=\"LEGACY_UI2016_subCategoryNavi\">\n".trim($res))."\n</div>\n";
    return "<a href=\"/\">トップへ</a><div class=\"newsArticle\">".$res."</div>";
  }

  /* 記事だけが残るまで消す */
  function rmbloat (string $body): string {
    $res = preg_replace('/<!DOCTYPE html>(.*?)<div class="newsArticle">/s', "", $body);
    $res = preg_replace('/<!--\/newsArticle-->(.*?)<\/html>/s', "", $res);
    $res = preg_replace('/<p class="reactions">(.*?)<\/p>/s', "", $res);
    $res = preg_replace('/<ul class="diaryUtility\d*">(.*?)<\/ul>/s', "", $res);
    $res = preg_replace('/<table>(.*?)<\/table>/s', "", $res);
    $res = preg_replace('/<div class="adsense0(.*?)<\/div>/s', "", $res);
    $res = preg_replace('/<div class="adsense(.*?)<\/div>/s', "", $res);
    $res = preg_replace('/<div class="subInfo">(.*?)<\/div>/s', "", $res);
    $res = preg_replace("/(^[\r\n]*|[\r\n]+)[\s\t]*[\r\n]+/", "\n", $res);
    $res = str_replace("<!--article_image-->", "", $res);
    $res = str_replace("<!--/article_image-->", "", $res);
    $res = preg_replace("/<!--(.*)-->/", "", $res);
    $res = str_replace("<!--", "", $res);
    $res = str_replace("https://news-image.mixi.net", IMGPROXY."/news-image.mixi.net", $res);
    $res = trim("<div class=\"newsArticle\">\n".trim($res))."\n    </div>\n";
    return "<a href=\"/\">トップへ</a>".$res;
  }

  function getfront (): string {
    return '<div class="newsArticle"><div class="newsCategoryList"><div class="heading08"><h1>ニュースカテゴリ一覧</h1></div><ul class="newsCategoryList"><li class="newCategoryList"><a href="'.DOMAIN.'/?url=list_news_category.pl?id=7">エンタメ</a></li><li class="newCategoryList"><a href="'.DOMAIN.'/?url=list_news_category.pl?id=3">トレンド</a></li><li class="newCategoryList"><a href="'.DOMAIN.'/?url=list_news_category.pl?id=1">社会</a></li><li class="newCategoryList"><a href="'.DOMAIN.'/?url=list_news_category.pl?id=4">地域</a></li><li class="newCategoryList"><a href="'.DOMAIN.'/?url=list_news_category.pl?id=9">ゲーム・アニメ</a></li><li class="newCategoryList"><a href="'.DOMAIN.'/?url=list_news_category.pl?id=8">IT・インターネット</a></li><li class="newCategoryList"><a href="'.DOMAIN.'/?url=list_news_category.pl?id=6">スポーツ</a></li><li class="newCategoryList"><a href="'.DOMAIN.'/?url=list_news_category.pl?id=5">海外</a></li><li class="newCategoryList"><a href="'.DOMAIN.'/?url=list_news_category.pl?id=10">コラム</a></li><li class="newCategoryList"><a href="'.DOMAIN.'/?url=list_news_category.pl?id=2">ライフスタイル</a></li></ul></div></div>';
  }

  /* 記事の受取 */
  function get (string $url = ""): array {
    $res = [];
    $ch = curl_init();
    curl_setopt($ch, CURLOPT_URL, "https://news.mixi.jp/".$url);
    curl_setopt($ch, CURLOPT_HTTPHEADER, [
      "HTTP/1.0",
      "Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
      "Accept-Language: en-US,en;q=0.5",
      "Connection: keep-alive",
      "User-Agent: Mozilla/5.0 (X11; Linux x86_64; rv:68.0) Gecko/20100101 Firefox/68.0"
    ]);
    curl_setopt($ch, CURLOPT_RETURNTRANSFER, TRUE);
    $body = mb_convert_encoding(curl_exec($ch), "UTF-8", "EUC-JP");
    $httpCode = curl_getinfo($ch, CURLINFO_HTTP_CODE);
    curl_close($ch);

    if ($body && $httpCode == 200) {
      $res["title"] = gettitle($body);
      $res["content"] = isarticle($url) ? rmbloat($body) : rmcbloat($body);
      if (isarticle($url)) $res["img"] = getimg($body);
      $res["desc"] = getdesc($res["content"]);
    }

    return $res;
  }

  // URLを受け取って、mixiにわかれるURLに変更
  $gurl = str_replace("/?url=", "", htmlspecialchars($_SERVER["REQUEST_URI"]));

  // デフォルト＝エラー
  $out = ["title" => "見つけられない", "content" => '<div class="newsArticle"><div class="articleHeading02"><div class="headingArea"><h1>見つけられなかった</h1></div></div><div class="contents clearfix"><div class="article decoratable"><p>ごめんね！</p></div></div>'];

  // $gurlは「/」だったら、トップページを表示する。記事だったら、記事を表示する。
  if ($gurl == "/") $out = ["title" => "トップページ", "content" => getfront()];
  else if (isarticle($gurl)) $out = get($gurl);
  else if (iscategory($gurl)) $out = get($gurl);
?>
<!DOCTYPE html>
<html lang="ja">
  <head>
    <meta content="text/html; charset=euc-jp" http-equiv="content-type" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <?php if (isset($out["desc"])) { ?>
    <meta property="og:title" content="spliti 〜 <?= $out["title"] ?>" />
    <meta property="og:type" content="article" />
    <meta property="og:description" content="<?= $out["desc"] ?>" />
    <meta property="og:url" content="<?= $gurl ?>" />
    <?php } ?>
    <?php if (isset($out["img"])) { ?><meta name="thumbnail" content="<?= $out["img"] ?>" /><?php } ?>
    <title>spliti 〜 <?= $out["title"] ?></title>
    <link rel="stylesheet" type="text/css" href="/style.css" />
  </head>
  <body>
    <?= $out["content"] ?>
    <p class="footer">
      Spliti 1.1.0 |
      <a href="https://gitler.moe/suwako/spliti"><img src="/git.png" alt="Git"></a> |
      <a href="https://076.moe/">０７６</a>
    </p>
  </body>
</html>
