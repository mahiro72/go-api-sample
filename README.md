# Go API Sample

## 内容

シンプルな掲示板サービス

### 機能
- ユーザー登録
- ユーザー取得
- ユーザーの投稿作成
  - Basic認証が必要
- ユーザーの投稿一覧取得


## 動作確認
ユーザー登録

```
curl -X POST http://localhost:8080/users --data '{"name":"hoge"}'
```

```json
{"id":"96b4bad9-fb9c-4cd0-b518-b9529b0112f8","name":"hoge","created_at":"2023-10-21T15:56:05.968516+09:00"}
```

<br>

ユーザー取得

```
curl http://localhost:8080/users/{userID}
```

```json
{"id":"96b4bad9-fb9c-4cd0-b518-b9529b0112f8","name":"hoge","created_at":"2023-10-21T15:56:05.968516+09:00"}
```

<br>

ユーザーの投稿作成

Basic認証を使う

```
curl -X POST http://localhost:8080/users/1a72c3d6-83c7-47f8-82fd-cba3415d53f4/posts \
--header 'Authorization: Basic xxxx' \
--data '{"title":"hello","content":"hoge"}'
```

```json
{"id":"cf7fe1cc-0c9b-4bdf-9b7b-102351a3c338","title":"hello","content":"hoge","created_at":"2023-10-21T16:04:48.219667+09:00"}
```

<br>

ユーザーの投稿一覧取得

```
curl http://localhost:8080/users/1a72c3d6-83c7-47f8-82fd-cba3415d53f4/posts/list
```

```json
[{"id":"03beabab-e6c3-43c5-97eb-2f7362f18b98","title":"hello","content":"hoge","created_at":"2023-10-21T16:04:41.74259+09:00"},{"id":"4d0055f5-bfd0-445e-b3c2-d5de00ecea3e","title":"hello","content":"hoge","created_at":"2023-10-21T16:04:46.515345+09:00"}]
```