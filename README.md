# Go API Sample

## 内容

ユーザーのタスク管理アプリ

### 登場人物
- ユーザー
- タスク
  - ユーザーは複数タスクを持てる

### 機能
- ユーザー登録
- ユーザー取得
- ユーザーの投稿登録
- ユーザーの投稿取得
- なりすましOK


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

投稿する

Basic認証を使う

```
curl -X POST http://localhost:8080/users/1a72c3d6-83c7-47f8-82fd-cba3415d53f4/posts \
--header 'Authorization: Basic xxxx' \
--data '{"title":"hello","content":"hoge"}'
```

<br>

ユーザーの投稿一覧を取得する

```
curl http://localhost:8080/users/1a72c3d6-83c7-47f8-82fd-cba3415d53f4/posts/list
```

