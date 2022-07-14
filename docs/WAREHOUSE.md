# Warehouse
## Route api/v1

#### [GET] /warehouse
```json
{
    "code": 200,
    "data": [
      {
        "id": 1,
        "address": "rua 1",
        "telephone": "11",
        "warehouse_code": "1",
        "minimum_capacity": 1,
        "minimum_temperature": 2,
        "locality_id": 1
      }
    ]
}
```
#### [GET] /warehouse/1
```json
{
    "code": 200,
    "data": {
        "id": 1,
        "address": "rua 1",
        "telephone": "11",
        "warehouse_code": "1",
        "minimum_capacity": 1,
        "minimum_temperature": 2,
        "locality_id": 1
      }    
}
```
#### [POST] /warehouse 
```json
{
    "address": "rua 4",
    "telephone": "12345678",
    "warehouse_code": "1212",
    "minimum_capacity": 2,
    "minimum_temperature": 22,
    "locality_id": 1
}
```

#### [PATCH] /warehouse/1
```json
{
    "address": "rua 4",
    "warehouse_code": "1212",
    "minimum_capacity": 2,
    "minimum_temperature": 22,
    "locality_id": 1
}
```

#### [DELETE] /warehouse/1
```shell
  status 204 no content
```
