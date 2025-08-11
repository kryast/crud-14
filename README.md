CRUD ke-14

POST
curl -X POST http://localhost:8080/product \
-H "Content-Type: application/json" \
-d '{
  "name": "Laptop Gaming",
  "price": 18500000.50,
  "stock": 10
}