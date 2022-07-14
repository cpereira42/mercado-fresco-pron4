# Employee

## Route api/v1

#### [GET] /employees
```json
{
    "code": 200,
    "data": [
      {
        "id": 1,
        "card_number_id": "1",
        "first_name": "mercado",
        "last_name": "livre",
        "warehouse_id": 1
      }
    ]
}
```
#### [GET] /employees/1
```json
{
    "code": 200,
    "data": {
        "id": 1,
        "card_number_id": "1",
        "first_name": "mercado",
        "last_name": "livre",
        "warehouse_id": 1
      }    
}
```
#### [POST] /employees 
```json
{
    "card_number_id": "2",
    "first_name": "mercado",
    "last_name": "livre",
    "warehouse_id": 1
}
```

#### [PATCH] /employees/1
```json
{
    "card_number_id": "2",
    "first_name": "mercado",
    "warehouse_id": 1
}
```

#### [DELETE] /employees/1
```shell
  status 204 no content
```
