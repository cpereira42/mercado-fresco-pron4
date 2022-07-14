# Locality
## Route api/v1

#### [GET] /localities/reportSellers
```json
{
    "code": 200,
    "data": [
      {
        "locality_id": 1,
        "locality_name": "São Paulo",
        "sellers_count": 2
      }
    ]
}
```
#### [GET] /localities/reportSellers?id=1
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
#### [POST] /localities
```json
{
    "id": 44445,
    "locality_name": "London",
    "province_name": "No se",
    "country_name": "England"
}
```

