# Report Products
## Route api/v1

#### [GET] /products/reportRecords
```json
{
    "code": 200,
    "data": [
      {
        "product_id": 1,
        "records_count": 1,
        "description": "teste"
      }
    ]
}
```
#### [GET] /products/reportRecords?id=1
```json
{
    "code": 200,
    "data": {
        "product_id": 1,
        "records_count": 1,
        "description": "teste"
      }    
}
```
#### [POST] /productsRecords
```json
{
    "purchase_price": 10,
    "sale_price": 15,
    "product_id": 1
}
```

