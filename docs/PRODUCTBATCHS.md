# Product Batchs
## Route api/v1

#### [GET] /sections/reportProducts
```json
{
    "code": 200,
    "data": [
      {
        "section_id": 1,
        "section_number": 1,
        "products_count": 1
      }
    ]
}
```
#### [GET] /sections/reportProducts?id=1
```json
{
    "code": 200,
    "data": {
        "section_id": 1,
        "section_number": 1,
        "products_count": 1
      }    
}
```
#### [POST] /productBatches
```json
{
	"batch_number": "56",
	"current_quantity": 200,
	"current_temperature": 20,
	"due_date": "2022-04-04",
	"initial_quantity": 10,
	"manufacturing_date": "2020-04-04 14:30:10",
	"manufacturing_hour": "2020-05-01 14:20:15",
	"minimum_temperature":5,
	"product_id": 1,
	"section_id": 6
}
```

