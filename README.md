# michisirube_camp_backend
## API LIST
### recommend API
スポットをリコメンドするAPI
### evaluation API
スポットを評価するAPI

## Request List
### スポットをリコメンドする
HTTPメソッド： `POST`  
リクエスト
```
{
  "emotion": {
    "value": "int",
  },
  "location": {
    "latitude": "float",
    "longitude": "float"
  }
  "time": "int"
}
```
レスポンス
```
{
  "location": {
    "id": "int",
    "name": "string",
    "latitude": "float",
    "longitude": "float"
  },
  "time": "int"
}
```
### スポットを評価する
HTTPメソッド: `POST`  
リクエスト
```
{
  "location": {
    "id": "int",
    "evaluation": {
      "emotion": {
        "value": "int",
      },
      "value": "int"
    }
  }
}
```
レスポンス
```
{
  "status": "bool"
}
```
## Table LIST
### Locate table
|Key|Value|
|:-|:-:|
|id|int|
|locate|string|
|latitude|float|
|longitude|float|
### Emotional table
|Key|Value|
|:-|:-:|
|id|int|
|glad|int|
|anger|int|
|sad|int|
|pleasant|int|
### Log table
|Key|Value|
|:-|:-:|
|id|int|
|count|int|
|timestamp|string|
|emotion|int|
|value|int|
