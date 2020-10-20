gRPC 動手玩
===

# 環境
- 使用 127.0.0.1:9999
- go 執行在本機
- php 執行在 docker container 內

# commands

## go
```
protoc greeting.proto --go_out=.

protoc greeting.proto --go_out=plugins=grpc:.
```

## php
```
protoc greeting.proto --php_out=php 

# client
protoc greeting.proto --grpc_out=php --plugin=protoc-gen-grpc=./grpc_php_plugin 

# server
protoc greeting.proto --grpc_out=php --plugin=protoc-gen-grpc=./protoc-gen-php-grpc 
```

# references
[官方文件](https://grpc.io/docs/what-is-grpc/)

[gRPC ，使用 PHP 實作](https://medium.com/@poyu677/grpc-%E4%BD%BF%E7%94%A8-php-%E5%AF%A6%E4%BD%9C-ab485e9f1044)

[spiral/php-grpc](https://github.com/spiral/php-grpc)
