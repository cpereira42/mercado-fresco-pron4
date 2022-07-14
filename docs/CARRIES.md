# Carries
## Route api/v1

#### [GET] /localities/reportCarries
```json
{
    "code": 200,
    "data": [
      {
        "locality_id": 1,
        "locality_name": "São Paulo",
        "carries_count": 5
      }
    ]
}
```
#### [GET] /localities?id=1
```json
{
    "code": 200,
    "data": {
        "locality_id": 1,
        "locality_name": "São Paulo",
        "sellers_count": 2
      }    
}
```
#### [POST] /carries 
```json
{
    "cid": "WX-20", 
    "company_name": "some name", 
    "address": "corrientes 800", 
    "telephone": "4567-4567", 
    "locality_id": 1
}
```

