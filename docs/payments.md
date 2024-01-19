# Payment API Spec

## Create Payment API

Endpoint : POST /api/payments


Request Body :

```json
{
    "name": "Cash",
    "type": "CASH"
}
```

Response Body Success :

```json
{
  "success": true,
  "message": "Success",
  "data": {
    "id": 2,
    "name": "Cash",
    "type": "CASH"
  }
}
```

Response Body Error :

```json
{
    "errors" : "name is required"
}
```

## Update Payment API

Endpoint : PUT /api/payments/:id

Request Body :

```json
{
    "name": "Cash",
    "type": "CASH"
}
```

Response Body Success :

```json
{
  "success": true,
  "message": "Success",
  "data": {
    "id": 1,
    "name": "Cashssss",
    "type": "CASH"
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
    "payments": [
      {
        "id": 1,
        "name": "Cash",
        "type": "CASH"
      },
      {
        "id": 2,
        "name": "Cash",
        "type": "CASH"
      }
    ],
    "meta": {
      "total": 2,
      "limit": 10,
      "skip": 0
    }
  }
}
```

Response Body Error :

```json
{
    "errors" : "name is not found"
}
```

## Get Payment API

Endpoint : GET /api/payments/:id

Headers :
- Authorization : token

Response Body Success :

```json
{
  "success": true,
  "message": "Success",
  "data": {
    "id": 1,
    "name": "Cash",
    "type": "CASH"
  }
}
```

Response Body Error :

```json
{
    "errors" : "name is not found"
}
```

## Remove Payment API

Endpoint : DELETE /api/payments/:id

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