# Seller
## Route api/v1

#### [GET] /sellers
```json
{
    "code": 200,
    "data": [
      {
        "id": 1,
        "cid": "cid1",
        "company_name": "Mercado",
        "address": "rua 1",
        "telephone": "111",
        "locality_id": 1
      }
    ]
}
```
#### [GET] /sellers/1
```json
{
    "code": 200,
    "data": {
        "id": 1,
        "cid": "cid1",
        "company_name": "Mercado",
        "address": "rua 1",
        "telephone": "111",
        "locality_id": 1
      }    
}
```
#### [POST] /sellers 
```json
{
    "cid": "cid2",
    "company_name": "Mercado",
    "address": "rua 1",
    "telephone": "111",
    "locality_id": 1
}
```

#### [PATCH] /sellers/1
```json
{
    "cid": "cid2",
    "company_name": "Mercado",
    "telephone": "111",
    "locality_id": 1
}
```

#### [DELETE] /sellers/1
```shell
  status 204 no content
```
