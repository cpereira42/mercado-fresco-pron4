# Buyer

## Route api/v1

#### [GET] /buyers
```json
{
    "code": 200,
    "data": [
      {
        "id": 1,
        "card_number_id": "123456",
        "first_name": "Mercado",
        "last_name": "Livre"
      }
    ]
}
```
#### [GET] /buyers/1
```json
{
    "code": 200,
    "data": {
        "id": 1,
        "card_number_id": "123456",
        "first_name": "Mercado",
        "last_name": "Livre"
      }    
}
```
#### [POST] /buyers 
```json
{
    "card_number_id": "402",
    "first_name": "Joe",
    "last_name": "Doe"
}
```

#### [PATCH] /buyers/1
```json
{
    "card_number_id": "402",
    "last_name": "Doe"
}
```

#### [DELETE] /buyers/1
```shell
  status 204 no content
```

