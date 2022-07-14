# Section
## Route api/v1

#### [GET] /sections
```json
{
    "code": 200,
    "data": [
      {
        "id": 1,
        "section_number": 1,
        "current_capacity": 1,
        "current_temperature": 1,
        "maximum_capacity": 1,
        "minimum_capacity": 1,
        "minimum_temperature": 1,
        "warehouse_id": 1,
        "product_type_id": 1
      }
    ]
}
```
#### [GET] /sections/1
```json
{
    "code": 200,
    "data": {
        "id": 1,
        "section_number": 1,
        "current_capacity": 1,
        "current_temperature": 1,
        "maximum_capacity": 1,
        "minimum_capacity": 1,
        "minimum_temperature": 1,
        "warehouse_id": 1,
        "product_type_id": 1
      }    
}
```
#### [POST] /sections 
```json
{
    "section_number": 8,
    "current_temperature": 795,
    "minimum_temperature": 3,
    "current_capacity": 15,
    "minimum_capacity": 23,
    "maximum_capacity": 456,
    "warehouse_id": 1,
    "product_type_id": 1
}
```

#### [PATCH] /sections/1
```json
{
    "section_number": 8,
    "current_temperature": 795,
    "minimum_capacity": 23,
    "maximum_capacity": 456,
    "warehouse_id": 1,
    "product_type_id": 1
}
```

#### [DELETE] /sections/1
```shell
  status 204 no content
```
