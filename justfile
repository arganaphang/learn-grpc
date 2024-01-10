run-api:
  cd api ; go run ./cmd

gen-api:
  cd api ; protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    ./user/user.proto

run-server:
  cd server ; python main.py

gen-server:
  cd server ; python -m grpc_tools.protoc -I./user --python_out=./user --pyi_out=./user --grpc_python_out=./user ./user/user.proto
