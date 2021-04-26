# TownPost - Server

## 開発環境の起動

go + realizeを使ってホットリロードに対応しています。

```shell
# 起動
$ make dev-up

# 開発環境の削除
$ make dev-down

# ログ
$ make dev-log-(api|db)
```

## code generate

### oapi codegen

[oapi-codegen](https://github.com/deepmap/oapi-codegen)より、[go-chi/chi](https://github.com/go-chi/chi)を用いたコードをopenAPI定義から自動生成します。

```shell
$ make gen-api
```

### mockgen

```shell
$ make gen-mock
```

### generate shcema by ent

```shell
$ make gen-schema
```

## document

OpenAPIの定義からSwaggerUIを作成します。

```shell
$ make doc
```

## mock server

8080ポートで起動するため、開発環境を開いている際は、`$ make dev-stop`などで一度開発環境を止めるか、モックを別ポートで起動してください。

```shell
# mockを作成
$ make build-mock

# run
$ cd mock-server && npm start
```

```shell
# request sample
$ curl --request POST \
  --url http://localhost:8080/signup \
  --header 'Content-Type: application/json' \
  --data '{
	
}'
```

## 環境変数

```
ENV = dev | prd
GCP_PROJECT_ID = string
GCP_SECRET_ID = string
```

ENVが`prd`の時は、gcpの本番環境に接続しに行きます。基本的には、devで開発を行ってください。