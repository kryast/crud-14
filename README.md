CRUD ke-14

POST
curl -X POST http://localhost:8080/product \
-H "Content-Type: application/json" \
-d '{
  "name": "Laptop Gaming",
  "price": 18500000.50,
  "stock": 10
}

GET
curl -X GET http://localhost:8080/product
curl -X GET http://localhost:8080/product/1

PUT
curl -X PUT http://localhost:8080/product/1 \
-H "Content-Type: application/json" \
-d '{
  "name": "Laptop Gaming Update",
  "price": 19000000.00,
  "stock": 8
}'

DELETE
curl -X DELETE http://localhost:8080/product/1