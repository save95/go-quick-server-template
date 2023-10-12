# go-quick-server-template
Golang 服务端模版

## 目录说明

## Usage

```shell
git clone --depth=1 -b main https://github.com/save95/go-quick-server-template.git server-api
#git clone --depth=1 -b develop https://github.com/save95/go-quick-server-template.git server-api
#git clone --depth=1 -b feature/v1.0-captcha-dev https://github.com/save95/go-quick-server-template.git server-api

cd server-api
rm -rf .git

git init --initial-branch=main
git remote add origin https://xxxxxx
git add .
git commit -m "init"
git push -u origin main
```


## 编译

因使用了 `https://github.com/mattn/go-sqlite3` 包需要启用 CGO 编译

```shell
# https://github.com/FiloSottile/homebrew-musl-cross

brew install filosottile/musl-cross/musl-cross

```

```shell
make build

```

## 启动

```shell

./main
# 等同于 ./main -config=config/config.toml -mode=all

# 指定配置和应用模块
./main -config=https://www.domain.com/app/config.toml -mode=web

# 执行一次性脚本命令
./main -config=https://www.domain.com/app/config.toml -mode=cmd cmd.name=example-simple

```

