# User API Spec

## Create User API

Endpoint : POST /api/users

Request Body :

```json
{
    "name": "Anisa Subandono",
    "email":"anisa@gmail.com",
    "password":"123456"
}
```

Response Body Success :

```json
{
    "success": true,
    "message": "Success",
    "data": {
        "id": 3,
        "name": "Anisa Subandono",
        "email": "anisa@gmail.com",
        "created_at": "2023-02-27T22:14:03.471+07:00",
        "updated_at": "2023-02-27T22:14:03.471+07:00"
  }
}
```

Response Body Error :

```json
{
    "errors" : "email already registered"
}
```

## Update User API

Endpoint : PUT /api/users/:id

Headers :
- Authorization : token

Request Body :

```json
{
    "name": "Anisa nis nis",
    "email":"anisa@gmail.com"
}
```

Response Body Success :

```json
{
  "success": true,
  "message": "Success",
  "data": {
    "id": 3,
    "name": "Anisa nis nis",
    "email": "anisa@gmail.com",
    "created_at": "2023-02-27T15:14:03.471Z",
    "updated_at": "2023-02-27T15:27:05.943Z"
  }
}
```

Response Body Error :

```json
{
    "errors" : "Name length max 100"
}
```

## Get All User API

Endpoint : GET /api/users

Headers :
- Authorization : token

Response Body Success :

```json
{
  "success": true,
  "message": "Success",
  "data": {
    "meta": {
      "total": 3,
      "limit": 0,
      "skip": 0
    },
    "users": [
      {
        "id": 1,
        "name": "Anisa Wulandari",
        "email": "wulandari@gmail.com",
        "created_at": "2023-02-27T15:12:18.095Z",
        "updated_at": "2023-02-27T15:12:18.095Z"
      },
    ]
  }
}
```

Response Body Error :

```json
{
    "errors" : "email is not found"
}
```

## Get User API

Endpoint : GET /api/users/:id

Headers :
- Authorization : token

Response Body Success :

```json
{
  "success": true,
  "message": "Success",
  "data": {
    "id": 3,
    "name": "Anisa nis nis",
    "email": "anisa@gmail.com",
    "created_at": "2023-02-27T15:14:03.471Z",
    "updated_at": "2023-02-27T15:27:05.943Z"
  }
}
```

Response Body Error :

```json
{
    "errors" : "email is not found"
}
```

## Remove User API

Endpoint : DELETE /api/users/:id

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
    "errors" : "email is not found"
}
```