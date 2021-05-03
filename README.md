# golang server template

アーキテクチャ、使っている環境については、wikiに書いておきます。

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

## mockgen

```shell
$ make gen-mock
```

## 環境変数について

本環境では以下環境変数に対応しています。

- ENV
    - 開発環境のコントロール
    - development | production
- SESSION_COOKIE_NAME
    - セッションを持つクッキーの名前