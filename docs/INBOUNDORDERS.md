# Inbound Orders
## Route api/v1

#### [GET] /employees/reportInboundOrders
```json
{
    "code": 200,
    "data": [
      {
        "id": 2,
        "card_number_id": "402323",
        "first_name": "Jhon",
        "last_name": "Jhon",
        "warehouse_id": 1,
        "inbound_orders_count": 0
      }
    ]
}
```
#### [GET] /employees/reportInboundOrders=2
```json
{
    "code": 200,
    "data": {
            "id": 2,
            "card_number_id": "402323",
            "first_name": "Jhon",
            "last_name": "Jhon",
            "warehouse_id": 1,
            "inbound_orders_count": 5  
      }    
}
```
#### [POST] /inboundOrders
```json
{
    "order_number": "order#1",
    "employee_id":4,
    "product_batch_id": 1,
    "warehouse_id": 1
}
```

