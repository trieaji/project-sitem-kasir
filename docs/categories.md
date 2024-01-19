# Category API Spec

## Create Category API

Endpoint : POST /api/categories

Headers :
- Authorization : token

Request Body :

```json
{
    "name": "Electronics"
}
```

Response Body Success :

```json
{
  "success": true,
  "message": "Success",
  "data": {
    "id": 2,
    "name": "Electronics"
  }
}
```

Response Body Error :

```json
{
    "errors" : "name is required"
}
```

## Update Category API

Endpoint : PUT /api/categories/:id

Headers :
- Authorization : token

Request Body :

```json
{
    "name": "Electronics"
}
```

Response Body Success :

```json
{
  "success": true,
  "message": "Success",
  "data": {
    "id": 2,
    "name": "Electronics"
  }
}
```

Response Body Error :

```json
{
    "errors" : "name is required"
}
```

## Get All Category API

Endpoint : GET /api/categories

Headers :
- Authorization : token

Response Body Success :

```json
{
  "success": true,
  "message": "Success",
  "data": {
    "meta": {
      "total": 1,
      "limit": 10,
      "skip": 0
    },
    "categories": [
      {
        "id": 1,
        "name": "Minuman"
      }
    ]
  }
}
```

Response Body Error :

```json
{
    "errors" : "name is not found"
}
```

## Get Category API

Endpoint : GET /api/categories/:id

Headers :
- Authorization : token

Response Body Success :

```json
{
  "success": true,
  "message": "Success",
  "data": {
    "id": 1,
    "name": "Minuman"
  }
}
```

Response Body Error :

```json
{
    "errors" : "name is not found"
}
```

## Remove Category API

Endpoint : DELETE /api/categories/:id

Headers :
- Authorization : token

Response Body Success :

```json
{
  "message": "Success",
  "success": true
}
```

Response Body Error :

```json
{
    "errors" : "name is not found"
}
```