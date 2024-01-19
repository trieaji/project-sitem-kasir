# Product API Spec

## Create Product API

Endpoint : POST /api/products

Headers :
- Authorization : token

Request Body :

```json
{
    "category_id": 2,
    "name": "Chiki ball ss",
    "image": "https://images.tokopedia.net/img/cache/500-square/product-1/2020/2/13/35604504/35604504_105cab00-a047-41a9-beaa-a27e755b61f2_1100_1100",
    "price": 7000,
    "stock": 100
}
```

Response Body Success :

```json
{
  "success": true,
  "message": "Success",
  "data": {
    "id": 3,
    "sku": "AAA",
    "name": "Chiki ball ss",
    "stock": 100,
    "price": 7000,
    "image": "https://images.tokopedia.net/img/cache/500-square/product-1/2020/2/13/35604504/35604504_105cab00-a047-41a9-beaa-a27e755b61f2_1100_1100",
    "category": null,
    "created_at": "2023-02-28T22:20:37.364+07:00",
    "updated_at": "2023-02-28T22:20:37.364+07:00"
  }
}
```

Response Body Error :

```json
{
    "errors" : "sku is required"
}
```

## Update Product API

Endpoint : PUT /api/products/:id

Headers :
- Authorization : token

Request Body :

```json
{
    "category_id": 2,
    "name": "Chiki ball",
    "image": "https://images.tokopedia.net/img/cache/500-square/product-1/2020/2/13/35604504/35604504_105cab00-a047-41a9-beaa-a27e755b61f2_1100_1100",
    "price": 7000,
    "stock": 100
}
```

Response Body Success :

```json
{
  "success": true,
  "message": "Success",
  "data": {
    "id": 1,
    "sku": "AAA",
    "name": "Chiki ball",
    "stock": 100,
    "price": 7000,
    "image": "https://images.tokopedia.net/img/cache/500-square/product-1/2020/2/13/35604504/35604504_105cab00-a047-41a9-beaa-a27e755b61f2_1100_1100",
    "category": null,
    "created_at": "2023-02-28T15:07:12.455Z",
    "updated_at": "2023-02-28T15:20:58.742Z"
  }
}
```

Response Body Error :

```json
{
    "errors" : "sku is required"
}
```

## Get All Product API

Endpoint : GET /api/products

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
    "products": [
      {
        "id": 2,
        "sku": "AAA",
        "name": "Chiki ball",
        "stock": 100,
        "price": 7000,
        "image": "https://images.tokopedia.net/img/cache/500-square/product-1/2020/2/13/35604504/35604504_105cab00-a047-41a9-beaa-a27e755b61f2_1100_1100",
        "category": {
          "id": 3,
          "name": "general"
        },
        "created_at": "2023-02-28T15:10:51.495Z",
        "updated_at": "2023-02-28T15:10:51.495Z"
      }
    ]
  }
}
```

Response Body Error :

```json
{
    "errors" : "sku is not found"
}
```

## Get Product API

Endpoint : GET /api/products/:id

Headers :
- Authorization : token

Response Body Success :

```json
{
  "success": true,
  "message": "Success",
  "data": {
    "id": 1,
    "sku": "AAA",
    "name": "Chiki ball",
    "stock": 100,
    "price": 7000,
    "image": "https://images.tokopedia.net/img/cache/500-square/product-1/2020/2/13/35604504/35604504_105cab00-a047-41a9-beaa-a27e755b61f2_1100_1100",
    "category": {
      "id": 3,
      "name": "general"
    },
    "created_at": "2023-02-28T15:07:12.455Z",
    "updated_at": "2023-02-28T15:07:12.455Z"
  }
}
```

Response Body Error :

```json
{
    "errors" : "sku is not found"
}
```

## Remove Product API

Endpoint : DELETE /api/products/:id

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
    "errors" : "sku is not found"
}
```