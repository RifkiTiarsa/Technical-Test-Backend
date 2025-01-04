# Technical Test - Backend Engineer

## Overview
This repository contains solutions to the technical test for the position of Back-End Engineer.

---

## Library
- godotenv  : `go get github.com/joho/godotenv`
- jwt       : `go get -u github.com/golang-jwt/jwt/v5`
- uuid      : `go get github.com/google/uuid`
- ORM       : `go get -u gorm.io/gorm`
- bycrypt   : `go get golang.org/x/crypto/bcrypt`

## Framework
- gin       : `go get github.com/gin-gonic/gin`

## Database
- postgreSQL : `go get -u gorm.io/driver/postgres`

---

## Task 1: API Design
### Features
1. **User Registration and Authentication**
2. **Viewing and Searching Products**
3. **Adding Items to a Shopping Cart**
4. **Completing a Purchase**

### RESTful Endpoints

| Feature                         | HTTP Method | Endpoint                  | Description                                                  | Auth Required |
|---------------------------------|-------------|---------------------------|--------------------------------------------------------------|---------------|
| User Registration              | POST        | `api/v1/auth/register`        | Register a new user                                           | No            |
| User Authentication (Login)    | POST        | `api/v1/auth/login`           | Authenticate a user and return a token                       | No            |
| View All Products              | GET         | `/api/v1/products`        | Fetch a list of all available products                       | No            |
| Search Products by Name        | GET         | `/api/v1/products?name` | Search products by name                                      | No            |
| Add Item to Cart               | POST        | `/api/v1/carts`            | Add a product to the user's shopping cart                    | Yes           |
| Checkout             | POST        | `/api/v1/checkout`        | Complete the purchase of items in the shopping cart          | Yes           |

### API Specification

### User Registration
**Request**  
- **Method**    : `POST`  
- **Endpoint**  : `/api/v1/auth/register`  
- **Headers**   :  
  - `Content-Type   : application/json`  
  - `Accept         : application/json`  
- **Body**          :
  ```json
    {
        "username": "dummy",
        "email": "dummy3@gmail.com",
        "password": "Rahasia08@"
    }

**Response**
-  **Status**    : `201 Created`
-  **Body**      :
   ```json
     {
        "status": {
            "code": 201,
            "message": "Register user successfully"
        },
        "data": {
            "id": "b212abbc-e52d-4752-bb48-0a2a39ce102d",
            "username": "dummy",
            "email": "dummy3@gmail.com",
            "created_at": "2025-01-04T23:02:25.942461532+07:00",
            "updated_at": "2025-01-04T23:02:25.942461532+07:00",
            "deleted_at": {
                "Time": "0001-01-01T00:00:00Z",
                "Valid": false
            }
        }
     }

### User Authentication
**Request**  
- **Method**    : `POST`  
- **Endpoint**  : `/api/v1/auth/login`  
- **Headers**   :  
  - `Content-Type   : application/json`  
  - `Accept         : application/json`  
- **Body**          :
  ```json
    {
        "email": "dummy3@gmail.com",
        "password": "Rahasia08@"
    }

**Response**
-  **Status**    : `200 OK`
-  **Body**      :
   ```json
     {
        "status": {
            "code": 200,
            "message": "Login successfully"
        },
        "data": {
            "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJyaWZraVRpYXJzYSIsImV4cCI6MTczNjAwOTA1OSwiaWF0IjoxNzM2MDA1NDU5LCJ1c2VyX2lkIjoiMjBhNGRhODQtNTRlNC00ZDNkLWE1YWMtNjFlZjI2Zjg4ZTRkIiwidXNlcm5hbWUiOiJUaWFyc2EiLCJlbWFpbCI6ImR1bW15MkBnbWFpbC5jb20ifQ.evq4c1VSp_IfXSe9_BZoSj4tnrJ9Cn1OrSvjq4nUWVM"
        }
     }

### Viewing All Products
**Request**  
- **Method**    : `GET`  
- **Endpoint**  : `/api/v1/products`
- **Headers**   :   
  - `Accept         : application/json`  

**Response**
-  **Status**    : `200 OK`
-  **Body**      :
   ```json
     {
        
        "status": {
        "code": 200,
        "message": "List product successfully"
        },
        "data": [
            {
                "id": "6d1f7864-cffe-4424-b369-f5db648f6634",
                "name": "Smartphone X1",
                "description": "Smartphone dengan layar 6.5 inci, RAM 4GB, dan penyimpanan 64GB",
                "category": {
                    "id": "0452e5f3-58a6-4a74-b9d7-851fa3d6d9d8",
                    "name": "Electronics"
                },
                "price": 4999000,
                "stock": 50,
                "rating": 4.5,
                "created_at": "2025-01-02T19:44:23.953461Z",
                "updated_at": "2025-01-02T19:44:23.953461Z",
                "deleted_at": {
                    "Time": "0001-01-01T00:00:00Z",
                    "Valid": false
                }
            },
            {
                "id": "29eb5a28-b023-4864-91a2-23cf1b8b718f",
                "name": "T-Shirt Basic",
                "description": "T-Shirt berbahan katun dengan berbagai pilihan warna",
                "category": {
                    "id": "2f7a42fa-457f-426d-a4b4-236d83a1e514",
                    "name": "Clothing"
                },
                "price": 150000,
                "stock": 96,
                "rating": 4.2,
                "created_at": "2025-01-02T19:44:23.953461Z",
                "updated_at": "2025-01-04T21:56:49.081037Z",
                "deleted_at": {
                    "Time": "0001-01-01T00:00:00Z",
                    "Valid": false
                }
            }
        ]
     }

### Search Products by Name
**Request**  
- **Method**    : `GET`  
- **Endpoint**  : `/api/v1/products?name=value`
- **Headers**   :   
  - `Accept         : application/json`  

**Response**
-  **Status**    : `200 OK`
-  **Body**      :
   ```json
     {
        "status": {
        "code": 200,
        "message": "Get product by name successfully"
        },
        "data": {
            "id": "29eb5a28-b023-4864-91a2-23cf1b8b718f",
            "name": "T-Shirt Basic",
            "description": "T-Shirt berbahan katun dengan berbagai pilihan warna",
            "category": {
                "id": "2f7a42fa-457f-426d-a4b4-236d83a1e514",
                "name": "Clothing"
            },
            "price": 150000,
            "stock": 96,
            "rating": 4.2,
            "created_at": "2025-01-02T19:44:23.953461Z",
            "updated_at": "2025-01-04T21:56:49.081037Z",
            "deleted_at": {
                "Time": "0001-01-01T00:00:00Z",
                "Valid": false
            }
        }
     }

### Add Item to Cart
**Request**  
- **Method**    : `POST`  
- **Endpoint**  : `/api/v1/carts` 
- **Headers**   :  
  - `Content-Type   : application/json`  
  - `Accept         : application/json`  
  - `Authorization  : Bearer Token`
- **Body**          :
  ```json
    {
        "user_id": "20a4da84-54e4-4d3d-a5ac-61ef26f88e4d",
        "product_id": "6d1f7864-cffe-4424-b369-f5db648f6634",
        "quantity": 1
    }

**Response**
-  **Status**    : `201 Created`
-  **Body**      :
   ```json
     {
        "status": {
        "code": 201,
        "message": "The product has been successfully added to the cart."
        },
        "data": {
            "id": "1c35369d-775a-4b38-a9cb-83c4fceab667",
            "user_id": "20a4da84-54e4-4d3d-a5ac-61ef26f88e4d",
            "product_id": "6d1f7864-cffe-4424-b369-f5db648f6634",
            "quantity": 1,
            "total_price": 4999000,
            "created_at": "2025-01-04T22:53:24.535331591+07:00",
            "updated_at": "2025-01-04T22:53:24.535331591+07:00",
            "deleted_at": {
                "Time": "0001-01-01T00:00:00Z",
                "Valid": false
            }
        }
     }

### Completing a Purchase/Checkout
**Request**  
- **Method**    : `POST`  
- **Endpoint**  : `/api/v1/checkout` 
- **Headers**   :  
  - `Content-Type   : application/json`  
  - `Accept         : application/json`  
  - `Authorization  : Bearer Token`
- **Body**          :
  ```json
    {
        "cart_id": "0073a119-8615-4fe7-a102-ef9fd93709fc",
        "payment_method": "BCA",
        "address": "Surabaya",
        "logistic_provider": "JNE"
    }

**Response**
-  **Status**    : `201 Created`
-  **Body**      :
   ```json
     {
        "status": {
        "code": 200,
        "message": "Checkout created successfully"
        },
        "data": {
            "id": "ad832b36-e1fc-4382-a17a-afd8a7f2b49a",
            "cart_id": "0073a119-8615-4fe7-a102-ef9fd93709fc",
            "amount": 75000,
            "payment_status": "pending",
            "payment_method": "BRI",
            "address": "Surabaya",
            "logistic_provider": "JNE",
            "shipping_status": "menunggu pembayaran",
            "created_at": "2025-01-04T23:29:59.929958726+07:00"
        }
     }

---

## Task 2: Indexing Strategy
### Queries and Index Recommendations

1. **Fetch a user by username**:
   - **Index**: Create an individual index on the `username` column.
   - Query: `SELECT * FROM Users WHERE username = 'value';`

2. **Fetch users who signed up after a certain date**:
   - **Index**: Create an individual index on the `created_at` column.
   - Query: `SELECT * FROM Users WHERE created_at > '2023-01-01';`

3. **Fetch a user by email**:
   - **Index**: Create an individual index on the `email` column.
   - Query: `SELECT * FROM Users WHERE email = 'value';

### Trade-offs
- **Read Performance**: Individual indexes improve query performance for read-heavy operations.
- **Write Performance**: Inserting or updating rows may take slightly longer due to index maintenance.
- **Composite Index**: Not recommended for these specific queries as the access patterns are independent.

---


