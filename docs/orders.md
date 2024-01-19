# Order API Spec

## Create Order API

Endpoint : POST /api/orders

Headers :
- Authorization : token

Request Body :

```json
{
    "payment_id": 2,
    "total_paid": 100000000,
    "products": [
        {
            "product_id": 2,
            "qty": 4
        },
        {
            "product_id": 4,
            "qty": 3
        }
    ]
}
```

Response Body Success :

```json
{
  "success": true,
  "message": "Success",
  "data": {
    "order": {
      "id": 6,
      "user_id": 1,
      "payment_type_id": 2,
      "total_price": 49000,
      "total_paid": 100000000,
      "total_return": 99951000,
      "receipt_id": "AAAAA",
      "products": [
        {
          "id": 1,
          "order_id": 6,
          "product_id": 2,
          "qty": 4,
          "price": 7000,
          "total_normal_price": 28000,
          "total_final_price": 28000,
          "created_at": "2023-02-28T22:48:53.261+07:00",
          "updated_at": "2023-02-28T22:48:53.261+07:00"
        },
        {
          "id": 2,
          "order_id": 6,
          "product_id": 4,
          "qty": 3,
          "price": 7000,
          "total_normal_price": 21000,
          "total_final_price": 21000,
          "created_at": "2023-02-28T22:48:53.261+07:00",
          "updated_at": "2023-02-28T22:48:53.261+07:00"
        }
      ],
      "payment_type": {
        "id": 0,
        "name": "",
        "type": "",
        "logo": ""
      },
      "updated_at": "2023-02-28T22:48:53.25+07:00",
      "created_at": "2023-02-28T22:48:53.25+07:00"
    },
    "products": [
      {
        "id": 1,
        "order_id": 6,
        "product_id": 2,
        "qty": 4,
        "price": 7000,
        "total_normal_price": 28000,
        "total_final_price": 28000,
        "created_at": "2023-02-28T22:48:53.261+07:00",
        "updated_at": "2023-02-28T22:48:53.261+07:00"
      },
      {
        "id": 2,
        "order_id": 6,
        "product_id": 4,
        "qty": 3,
        "price": 7000,
        "total_normal_price": 21000,
        "total_final_price": 21000,
        "created_at": "2023-02-28T22:48:53.261+07:00",
        "updated_at": "2023-02-28T22:48:53.261+07:00"
      }
    ]
  }
}
```

Response Body Error :

```json
{
    "errors" : "name is required"
}
```

## Get All Order API

Endpoint : GET /api/orders

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
    "orders": [
      {
        "id": 6,
        "user_id": 1,
        "payment_type_id": 2,
        "total_price": 49000,
        "total_paid": 100000000,
        "total_return": 99951000,
        "receipt_id": "AAAAA",
        "products": null,
        "payment_type": {
          "id": 2,
          "name": "Cash",
          "type": "CASH",
          "logo": ""
        },
        "updated_at": "2023-02-28T15:48:53.25Z",
        "created_at": "2023-02-28T15:48:53.25Z"
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

## Get Order API

Endpoint : GET /api/orders/:id

Headers :
- Authorization : token

Response Body Success :

```json
{
  "success": true,
  "message": "Success",
  "data": {
    "id": 6,
    "user_id": 1,
    "payment_type_id": 2,
    "total_price": 49000,
    "total_paid": 100000000,
    "total_return": 99951000,
    "receipt_id": "AAAAA",
    "products": [
      {
        "id": 1,
        "order_id": 6,
        "product_id": 2,
        "qty": 4,
        "price": 7000,
        "total_normal_price": 28000,
        "total_final_price": 28000,
        "product": {
          "id": 2,
          "sku": "AAA",
          "name": "Chiki ball",
          "stock": 100,
          "price": 7000,
          "image": "https://images.tokopedia.net/img/cache/500-square/product-1/2020/2/13/35604504/35604504_105cab00-a047-41a9-beaa-a27e755b61f2_1100_1100",
          "category": null,
          "created_at": "2023-02-28T15:10:51.495Z",
          "updated_at": "2023-02-28T15:10:51.495Z"
        },
        "created_at": "2023-02-28T15:48:53.261Z",
        "updated_at": "2023-02-28T15:48:53.261Z"
      },
      {
        "id": 2,
        "order_id": 6,
        "product_id": 4,
        "qty": 3,
        "price": 7000,
        "total_normal_price": 21000,
        "total_final_price": 21000,
        "product": {
          "id": 4,
          "sku": "AAA",
          "name": "Chiki ball ss",
          "stock": 100,
          "price": 7000,
          "image": "https://images.tokopedia.net/img/cache/500-square/product-1/2020/2/13/35604504/35604504_105cab00-a047-41a9-beaa-a27e755b61f2_1100_1100",
          "category": null,
          "created_at": "2023-02-28T15:44:18.36Z",
          "updated_at": "2023-02-28T15:44:18.36Z"
        },
        "created_at": "2023-02-28T15:48:53.261Z",
        "updated_at": "2023-02-28T15:48:53.261Z"
      }
    ],
    "payment_type": {
      "id": 2,
      "name": "Cash",
      "type": "CASH",
      "logo": ""
    },
    "updated_at": "2023-02-28T15:48:53.25Z",
    "created_at": "2023-02-28T15:48:53.25Z"
  }
}
```

Response Body Error :

```json
{
    "errors" : "name is not found"
}
```