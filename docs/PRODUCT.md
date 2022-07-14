# Warehouse
## Route api/v1

#### [GET] /warehouse
```json
{
    "code": 200,
    "data": [
      {
        "id": 1,
        "product_code": "product1",
        "description": "teste",
        "width": 15,
        "length": 15,
        "height": 15,
        "net_weight": 15,
        "expiration_rate": 15,
        "recommended_freezing_temperature": 12,
        "freezing_rate": 12,
        "product_type_id": 1,
        "seller_id": 1
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
        "product_code": "product1",
        "description": "teste",
        "width": 15,
        "length": 15,
        "height": 15,
        "net_weight": 15,
        "expiration_rate": 15,
        "recommended_freezing_temperature": 12,
        "freezing_rate": 12,
        "product_type_id": 1,
        "seller_id": 1
      }    
}
```
#### [POST] /warehouse 
```json
{
    "product_code": "prod2",
    "description": "celular",
    "width": 1.1,
    "length": 2.2,
    "height": 3.3,
    "net_weight": 4.4,
    "expiration_rate": 5.5,
    "recommended_freezing_temperature": 6.6,
    "freezing_rate": 7,
    "product_type_id": 8,
    "seller_id": 1 
}
```

#### [PATCH] /warehouse/1
```json
{
    "width": 1.1,
    "length": 2.2,
    "height": 3.3,
    "net_weight": 4.4,
    "expiration_rate": 5.5,
    "recommended_freezing_temperature": 6.6,
    "freezing_rate": 7,
    "product_type_id": 8,
    "seller_id": 1 
}
```

#### [DELETE] /warehouse/1
```shell
  status 204 no content
```
