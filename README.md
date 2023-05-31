# spliti

mixi向けプライバシーUI。

## 設置方法

### すべてのOS

```sh
$domain="example.com"
cd /var/www/htdocs
git clone https://gitler.moe/suwako/spliti.git && cd spliti
mv config.example.php config.php
find . -type f -name "config.php" -exec sed -i 's/mixi.owacon.moe/$domain/g'
```

### Linux

```sh
cp srv/nginx.conf /etc/nginx/sites-enabled/spliti.conf
find . -type f -name "/etc/nginx/sites-enabled/spliti.conf" -exec sed -i 's/ドメイン名/$domain/g'
/etc/init.d/nginx restart
```

### FreeBSD

```sh
cp srv/nginx.conf /usr/local/etc/nginx/sites-enabled/spliti.conf
find . -type f -name "/usr/local/etc/nginx/sites-enabled/spliti.conf" -exec sed -i 's/ドメイン名/$domain/g'
service nginx restart
```

### OpenBSD

```sh
cat /etc/httpd.conf srv/httpd.conf > /etc/httpd.conf
find . -type f -name "/etc/httpd.conf" -exec sed -i 's/ドメイン名/$domain/g'
rcctl restart httpd
```

## インスタンス一覧

### 一般ネット

| ウエブサイト | [クラフレ](http://jezf25zgvxlsvuzdzm6fg2hoetmruhy4uxnolyw46tuh4jugcwc7byqd.onion/Cloudflare%E3%82%92%E4%BD%BF%E3%82%8F%E3%81%AA%E3%81%84%E7%90%86%E7%94%B1%EF%BC%88%E3%83%AA%E3%83%81%E3%83%A3%E3%83%BC%E3%83%89%E3%83%BB%E3%82%B9%E3%83%88%E3%83%BC%E3%83%AB%E3%83%9E%E3%83%B3%EF%BC%89) | 注 |
| -- | -- | -- |
| [https://spliti.owacon.moe/](https://spliti.owacon.moe/) | 無 | 公式インスタンス |

### Tor

| オニオン | 注 |
| -- | -- |
| | |

### I2P

| イープサイト | 注 |
| -- | -- |
| | |
