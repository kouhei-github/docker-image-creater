## cobraの初期対応


```
$ docker compose build
$ docker compose up -d
# コンテナに接続
$ docker compose exec go ash

# コンテナ内で作業
$ go mod init nagamatsu-cobra
$ go mod tidy
$ go mod install
$ cobra-cli init
```

---
