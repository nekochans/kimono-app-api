# kimono-app-api
![ci-master](https://github.com/nekochans/kimono-app-api/workflows/ci-master/badge.svg)
[![Coverage Status](https://coveralls.io/repos/github/nekochans/kimono-app-api/badge.svg?branch=master)](https://coveralls.io/github/nekochans/kimono-app-api?branch=master)

着物アプリのバックエンドAPI（仮）

## 開発環境

- Docker

## ローカル実行

以下のスクリプトを実行して下さい。

`./docker-compose-up.sh`

[この記事](https://qiita.com/keitakn/items/f46347f871083356149b) のように `delve` を使ってデバックを行う場合は以下のスクリプトを実行して下さい。

`./docker-compose-up-debug.sh`

## Lintの実行

`docker-compose exec api sh` でアプリケーション用のコンテナに入ります。

`make lint` を実行して下さい。

もしくは `docker-compose exec api make lint` でも実行出来ます。

lintのルール等は以下を参考にして下さい。

https://golangci-lint.run/usage/linters/

ここで表示されたエラーは修正を行う必要があります。

一部のエラー内容は後で解説する `make format` コマンドでも修正可能です。

## ソースコードのフォーマット

`docker-compose exec api sh` でアプリケーション用のコンテナに入ります。

`make format` を実行して下さい。

もしくは `docker-compose exec api make format` でも実行出来ます。

このコマンドで自動修正されない物は自分で修正を行う必要があります。

## テストの実行

`docker-compose exec api sh` でアプリケーション用のコンテナに入ります。

`make test` を実行します。

もしくは `docker-compose exec api make test` でもテストを実行出来ます。
