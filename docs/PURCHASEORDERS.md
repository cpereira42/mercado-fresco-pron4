# Purchase Orders
## Route api/v1

#### [GET] /buyers/reportPurchaseOrders
```json
{
    "code": 200,
    "data": [
      {
        "id": 3,
        "card_number_id": "40232212",
        "first_name": "Omar",
        "last_name": "Barra",
        "purchase_orders_count": 2
      }
    ]
}
```
#### [GET] /buyers/reportPurchaseOrders?id=3
```json
{
    "code": 200,
    "data": {
        "id": 3,
        "card_number_id": "40232212",
        "first_name": "Omar",
        "last_name": "Barra",
        "purchase_orders_count": 2
      }    
}
```
#### [POST] /purchaseOrders
```json
{
    "purchase_price": 10,
    "sale_price": 15,
    "product_id": 1
}
```

