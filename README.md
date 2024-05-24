# Product API

This API is built using Go, Gin and Gorm. It provides CRUD operations for products.

## Endpoints

### GET /products

Returns a list of all products.

### GET /products/:id

Returns a single product by its ID.

### POST /products/create

Creates a new product. The request body should be a JSON object with the following properties:

- `name`: The name of the product (string)
- `description`: The description of the product (string)
- `price`: The price of the product (number)

### PUT /products/update/:id

Updates an existing product by its ID. The request body should be a JSON object with the following properties:

- `name`: The new name of the product (string)
- `description`: The new description of the product (string)
- `price`: The new price of the product (number)

### DELETE /products/delete/:id

Deletes a product by its ID.

## Error Handling

In case of an error, the API will return a JSON object with a `message` property containing a description of the error.
