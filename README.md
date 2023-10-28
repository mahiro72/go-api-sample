# Go API Sample

GoのAPIのサンプルリポジトリです。

気になる点や改善点などありましたらissueを立てたり、気軽に[@mahiro0x00](https://x.com/mahiro0x00)までご連絡ください。

## 内容

シンプルな掲示板サービス

<br>

※ セキュリティ周りやDB周りはふわっとしか書いてません。

※ アーキテクチャや責務の切り分けについてをメインに書いてるリポジトリです。


## アーキテクチャ構成
```
.
├── controller
│   ├── handler //APIエンドポイントに対応するハンドラー
│   ├── router //serverのエントリーポイント
│   └── view //response用の構造体に変換する
├── domain
│   ├── model //ドメインやドメインロジック
│   ├── repository //永続化のインターフェース
│   └── service //ドメインロジックのインターフェース
├── impl
│   ├── repository //domain/repositoryの具体的な実装
│   └── servic //domain/serviceの具体的な実装
├── main.go //このアプリケーションのエントリーポイント
├── persistence
│   └── inmemory
│       ├── database.go //inmemoryで永続化するロジック
│       └── dto //modelをinmemoryで扱うデータ構造に変換する
├── usecase //エンドポイントと1:1に対応するビジネスロジック
└── utils //便利関数など
```


## Server起動

```
go run main.go
```

Serverのhealthを確認
```
curl http://localhost:8080/health
```

response
```
health ok!
```

## 機能

機能は以下


- ユーザー登録
- ユーザー取得
- ユーザーの投稿作成
  - Basic認証が必要
- ユーザーの投稿一覧取得

<br>

### 機能の動作確認

ユーザー登録

```
curl -X POST http://localhost:8080/users --data '{"name":"hoge","password":"hoge"}'
```

response
```json
{"id":"96b4bad9-fb9c-4cd0-b518-b9529b0112f8","name":"hoge","created_at":"2023-10-21T15:56:05.968516+09:00"}
```

<br>

ユーザー取得

```
curl http://localhost:8080/users/{userID}
```

response
```json
{"id":"96b4bad9-fb9c-4cd0-b518-b9529b0112f8","name":"hoge","created_at":"2023-10-21T15:56:05.968516+09:00"}
```

<br>

ユーザーの投稿作成

Basic認証を使う

```
curl -X POST http://localhost:8080/users/{userID}/posts \
--header 'Authorization: Basic xxxx' \
--data '{"title":"hello","content":"hoge"}'
```

response
```json
{"id":"cf7fe1cc-0c9b-4bdf-9b7b-102351a3c338","title":"hello","content":"hoge","created_at":"2023-10-21T16:04:48.219667+09:00"}
```

<br>

ユーザーの投稿一覧取得

```
curl http://localhost:8080/users/{userID}/posts/list
```

response
```json
[{"id":"03beabab-e6c3-43c5-97eb-2f7362f18b98","title":"hello","content":"hoge","created_at":"2023-10-21T16:04:41.74259+09:00"},{"id":"4d0055f5-bfd0-445e-b3c2-d5de00ecea3e","title":"hello","content":"hoge","created_at":"2023-10-21T16:04:46.515345+09:00"}]
```
