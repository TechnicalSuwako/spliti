<?php
  include("../config.php");

  /* ページのタイトル */
  function gettitle (string $str): string {
    preg_match("/<title>(.*)<\/title>/", $str, $matches);
    return $matches[1];
  }

  function getimg (string $str): string {
    preg_match('/<img class="NEWS_tempPhoto__picture" src="(.*)" alt="">/', $str, $matches);
    $res = str_replace("https://", "https://imgproxy.owacon.moe/", $matches[1]);
    unset($matches);
    return $res;
  }

  function getdesc (string $str): string {
    $res = preg_replace('/<div class="newsArticle">(.*?)<\/div>/s', "", $str);
    return strip_tags($res);
  }

  /* 記事かの確認 */
  function isarticle (string $url): bool {
    $chk = explode("=", $url);
    if (isset($chk[0]) && $chk[0] == "view_news.pl?id" && isset($chk[1]) && isset($chk[2])) {
      $chk2 = explode("&amp;", $chk[1]);
      if (isset($chk2[1]) && $chk2[1] == "media_id") return true;
    }
    return false;
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
    $res = str_replace("news-image.mixi.net", "imgproxy.owacon.moe/news-image.mixi.net", $res);
    return trim("<div class=\"newsArticle\">\n".trim($res))."\n    </div>\n";
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
      $res["content"] = rmbloat($body);
      $res["img"] = getimg($body);
      $res["desc"] = getdesc($res["content"]);
    }

    return $res;
  }

  // URLを受け取って、mixiにわかれるURLに変更
  $gurl = str_replace("/?url=", "", htmlspecialchars($_SERVER["REQUEST_URI"]));

  // デフォルト＝エラー
  $out = ["title" => "見つけられない", "content" => '<div class="newsArticle"><div class="articleHeading02"><div class="headingArea"><h1>見つけられなかった</h1></div></div><div class="contents clearfix"><div class="article decoratable"><p>ごめんね！</p></div></div>'];

  // $gurlは「/」だったら、トップページを表示する。記事だったら、記事を表示する。
  if ($gurl == "/") $out = ["title" => "トップページ", "content" => '<div class="newsArticle"><div class="articleHeading02"><div class="headingArea"><h1>使い方</h1></div></div><div class="contents clearfix"><div class="article decoratable"><p><code>https://news.mixi.jp/view_news.pl?id=********&media_id=***</code>→<code>'.DOMAIN.'/?url=view_news.pl?id=********&media_id=***</code><br />例えば：<code>https://news.mixi.jp/view_news.pl?id=7327623&media_id=4</code>→<code>'.DOMAIN.'/?url=view_news.pl?id=7327047&media_id=262</code></p></div></div>'];
  else if (isarticle($gurl)) $out = get($gurl);
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
      <a href="https://gitler.moe/suwako/spliti"><img src="/git.png" alt="Git"></a> |
      <a href="https://076.moe/">０７６</a>
    </p>
  </body>
</html>
