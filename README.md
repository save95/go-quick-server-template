# go-quick-server-template
Golang 服务端模版

## 目录说明

## Usage

```shell
git clone --depth=1 -b main https://github.com/save95/go-quick-server-template.git server-api

cd server-api
rm -rf .git

git init --initial-branch=main
git remote add origin https://xxxxxx
git add .
git commit -m "init"
git push -u origin main
```

### swagger 

先使用以下命令安装 `swag` 命令

```shell
go get -u github.com/swaggo/swag/cmd/swag
```

- [go-swagger 注释文档](https://swaggo.github.io/swaggo.io/declarative_comments_format/api_operation.html)
- [go-swagger 注释中文文档](https://github.com/swaggo/swag/blob/master/README_zh-CN.md)

## 编译

因使用了 `https://github.com/mattn/go-sqlite3` 包需要启用 CGO 编译

```shell
# https://github.com/FiloSottile/homebrew-musl-cross

brew install filosottile/musl-cross/musl-cross

```

```shell
make build

```
