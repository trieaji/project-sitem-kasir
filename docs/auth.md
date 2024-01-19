# User API Spec

## Login User API

Endpoint : POST /api/users/login

Request Body :

```json
{
    "email": "anisa@gmail.com",
    "password": "123456"
}
```

Response Body Success :

```json
{
    "success": true,
    "message": "Success",
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE2NDYzNjk2NDQsInN1YiI6IjEifQ.OXOV-TjfCbCCJ7z1w1osQ1lz99rK89V_Ert_Y1JUfCM"
    }
}
```

Response Body Error :

```json
{
    "errors" : "email or password wrong"
}
```

## Logout User API

Endpoint : DELETE /api/users/logout

Headers :
- Authorization : token

Response Body Success :

```json
{
    "data" : "Oke kehapus"
}
```

Response Body Error :

```json
{
    "errors" : "Unauthorized"
}
```