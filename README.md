# medical-hackathon

### サーバー起動方法
``` 
$ cd medical-hackathon-api
docker-compose up -d
```

### API

#### エンドポイント

#### User
| コマンド | パス                         | 役割             |
| -------- | ---------------------------- | ---------------- |
| POST     | localhost:1323/api/user/new  | ユーザー新規作成 |
| GET      | localhost:1323/api/users     | ユーザー一覧取得 |
| GET      | localhost:1323/api/users/:id | ユーザー1件取得  |
| PUT      | localhost:1323/api/users/:id | ユーザー1件更新  |
| DELETE   | localhost:1323/api/users/:id | ユーザー1件削除  |

#### Clinic
| コマンド | パス                           | 役割               |
| -------- | ------------------------------ | ------------------ |
| POST     | localhost:1323/api/clinics/new | クリニック新規作成 |
| GET      | localhost:1323/api/clinics     | クリニック一覧取得 |
| GET      | localhost:1323/api/clinics/:id | クリニック1件取得  |
| PUT      | localhost:1323/api/clinics/:id | クリニック1件更新  |
| DELETE   | localhost:1323/api/clinics/:id | クリニック1件削除  |




