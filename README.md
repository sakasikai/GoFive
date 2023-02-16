# GoFive
douyin demo

### 1.Setup Basic Dependence
```shell
docker-compose up
```

### 2.Run User RPC Server
```shell
sh build.sh
./cmd/user/output/bin/UserService
```

### 2.Run API Server
```shell
go run ./cmd/api/main.go
```