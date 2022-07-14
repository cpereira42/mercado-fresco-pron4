# Purchase Orders
## Route api/v1

#### [GET] /buyers/reportPurchaseOrders
```json
{
  "code": 200,
  "data": {
    "id": 1,
    "order_date": "2008-11-11T13:23:44Z",
    "order_number": "1",
    "tracking_code": "1521",
    "buyer_id": 1,
    "product_record_id": 1,
    "order_status_id": 1
  }
```
#### [GET] /buyers/reportPurchaseOrders?id=3
```json
{
    "code": 200,
    "data": {
        "id": 3,
       "order_date": "2008-11-11T13:23:44Z",
       "order_number": "1",
       "tracking_code": "1521",
       "buyer_id": 1,
       "product_record_id": 1,
        "order_status_id": 1
      }    
}
```
#### [POST] /purchaseOrders
```json
{
    "order_date": "2021-04-04 13:23:44",
    "order_number": "1",
    "tracking_code": "123",
    "buyer_id": 1,
    "product_record_id": 1,
    "order_status_id": 1
}
```

