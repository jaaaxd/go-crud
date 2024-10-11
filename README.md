# GO CRUD APIs for Products and Users 

This repository provides a set of basic CRUD APIs built with Go, using the Gin and GORM framework. 

## Tech Stack

- **Go**
- **Gin** 
- **GORM & PostgreSQL** 
- **JWT Token & Bcrypt** 

## API Endpoints

### Products:
- **POST** `/products` - Create a new product
- **GET** `/products` - Get all products
- **GET** `/products/:id` - Get a product by ID
- **PUT** `/products/:id` - Update a product by ID
- **DELETE** `/products/:id` - Delete a product by ID

### Users:
- **POST** `/users/register` - Register a new user
- **POST** `/users/login` - User login and return a token
- **GET** `/users` - Get every user's details 
- **GET** `/users/:id` - Get user details by ID
- **PUT** `/users/:id` - Update user details by ID
- **DELETE** `/users/:id` - Delete a user by ID
