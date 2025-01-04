# Technical-Test

## Overview
This repository contains solutions to the technical test for the position of Back-End Engineer.

## Library
- godotenv  : go get github.com/joho/godotenv
- jwt       : go get -u github.com/golang-jwt/jwt/v5
- uuid      : go get github.com/google/uuid
- ORM       : go get -u gorm.io/gorm
- bycrypt   : go get golang.org/x/crypto/bcrypt

## Framework
- gin       : go get github.com/gin-gonic/gin

## Database
- postgreSQL : go get -u gorm.io/driver/postgres

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
| View Cart Items                | GET         | `/api/v1/carts`            | Fetch items in the user's shopping cart                      | Yes           |
| Complete Purchase              | POST        | `/api/v1/purchase`        | Complete the purchase of items in the shopping cart          | Yes           |

