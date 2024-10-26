#!/bin/bash

# Установите адрес и порт вашего gRPC сервера
SERVER="localhost:50051"

# Создание заказа
echo "Testing CreateOrder..."
grpcurl -d '{"overprice": 100, "Description": "Order description"}' \
  -plaintext $SERVER api.OrderService/CreateOrder

# Проверка статуса после создания заказа
echo -e "\nTesting GetOrder..."
grpcurl -d '{}' \
  -plaintext $SERVER api.OrderService/GetOrder

# Загрузка документа
echo -e "\nTesting UploadDocument..."
grpcurl -d '{
  "id": 1,
  "price": 200,
  "OverPrice": 150,
  "Description": "Document description",
  "courierlist": [{"id": 1, "name": "Courier 1", "type": "bike", "dist": 10}],
  "yourcourier": {"id": 2, "name": "Courier 2", "type": "car", "dist": 20}
}' \
  -plaintext $SERVER api.OrderService/UploadDocument

echo -e "\nAll tests completed."
