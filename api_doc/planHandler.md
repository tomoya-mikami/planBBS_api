# Plan API

使用する構造体
* [Plan](../struct/plan.md)

## プランを全取得

### Path
`/plan`

### Method
`GET`

### Request Body
なし

### Response Body
```
{
    "plan": [
        {
            "Title": "旅行プラン",
            "Description": "赤坂の旅行プランです",
            "Events": [
                {
                    "Title": "空港",
                    "Description": "羽田空港にむかいます",
                    "URL": "https://xxxx",
                    "Date": 1588491615665
                },
                {
                    "Title": "ねこ",
                    "Description": "ねこです",
                    "URL": "",
                    "Date": 1588491620000
                }
            ]
        }
    ]
}
```

## プランを保存する

### Path
/plan

### Method
POST

### Request Body
```
{
    "plan": [
        {
            "Title": "旅行プラン",
            "Description": "赤坂の旅行プランです",
            "Events": [
                {
                    "Title": "空港",
                    "Description": "羽田空港にむかいます",
                    "URL": "https://tokyo-haneda.com/index.html",
                    "Date": 1588491615665
                },
                {
                    "Title": "ねこ",
                    "Description": "ねこです",
                    "URL": "",
                    "Date": 1588491620000
                }
            ]
        }
    ]
}
```

### Response Body
```
{}
```

### 注意点
Dateのunixtimeはクライアント側で変換して送る
