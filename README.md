# spliti

mixi向けプライバシーUI。

## 設置方法

### すべてのOS

```sh
$domain="example.com"
cd /var/www/htdocs
git clone https://gitler.moe/suwako/spliti.git && cd spliti
find . -type f -name "config.json" -exec sed -i 's/mixi.076.moe/$domain/g'
```

### OpenBSD（オススメ）

```sh
make
doas make install
doas make config
cd /etc
wget https://076.moe/repo/webserver/relayd/spliti.conf
mv spliti.conf relayd.conf
find . -type f -name "/etc/relayd.conf" -exec sed -i 's/DOMAIN/$domain/g'
rcctl restart relayd
```

### Linux

**注意：BSD Makeをインストールして下さい。GNU Makeは未対応です。**

```sh
bmake
doas bmake install PREFIX=/usr
doas bmake config
cp /etc/nginx/sites-enabled
wget https://076.moe/repo/webserver/nginx/spliti.conf
find . -type f -name "/etc/nginx/sites-enabled/spliti.conf" -exec sed -i 's/DOMAIN/$domain/g'
/etc/init.d/nginx restart
```

### FreeBSD

```sh
make
doas make install
doas make config CNFPREFIX=/usr/local/etc
cp srv/nginx.conf /usr/local/etc/nginx/sites-enabled/spliti.conf
wget https://076.moe/repo/webserver/nginx/spliti.conf
find . -type f -name "/usr/local/etc/nginx/sites-enabled/spliti.conf" -exec sed -i 's/DOMAIN/$domain/g'
service nginx restart
```

### NetBSD

```sh
make
doas make install
doas make config CNFPREFIX=/usr/pkg/etc
cp srv/nginx.conf /usr/pkg/etc/nginx/sites-enabled/spliti.conf
wget https://076.moe/repo/webserver/nginx/spliti.conf
find . -type f -name "/usr/pkg/etc/nginx/sites-enabled/spliti.conf" -exec sed -i 's/DOMAIN/$domain/g'
service nginx restart
```

## インスタンス一覧

### 一般ネット

| ウエブサイト | [クラフレ](http://jezf25zgvxlsvuzdzm6fg2hoetmruhy4uxnolyw46tuh4jugcwc7byqd.onion/Cloudflare%E3%82%92%E4%BD%BF%E3%82%8F%E3%81%AA%E3%81%84%E7%90%86%E7%94%B1%EF%BC%88%E3%83%AA%E3%83%81%E3%83%A3%E3%83%BC%E3%83%89%E3%83%BB%E3%82%B9%E3%83%88%E3%83%BC%E3%83%AB%E3%83%9E%E3%83%B3%EF%BC%89) | 注 |
| -- | -- | -- |
| [https://mixi.076.moe/](https://mixi.076.moe/) | 無 | 公式インスタンス |

### Tor

| オニオン | 注 |
| -- | -- |
| | |

### I2P

| イープサイト | 注 |
| -- | -- |
| | |
