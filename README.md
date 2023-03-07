# spliti

mixi向けプライバシーUI。

## 設置方法

### すべてのOS

```sh
$domain="example.com"
cd /var/www/htdocs
git clone https://gitler.moe/TechnicalSuwako/spliti.git && cd spliti
mv config.example.php config.php
find . -type f -name "/etc/nginx/sites-enabled/spliti.conf" -exec sed -i 's/spliti.076.moe/$domain/g'
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
