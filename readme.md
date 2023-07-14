# Команды для установки grpc:
    go get -u github.com/golang/protobuf/protoc-gen-go
    go get google.golang.org/grpc

# Команда для генерации protobuf сущностей:
    protoc --go_out=../pb ./model/*.proto  

### 1. protoc - команда для запуска из env_path
### 2. --go_out= path, куда будет генерироваться {name}.pb.go
#### так как в proto-файлах указан go_package = "/model", то конечный путь будет: ..pb/model/
### 3. ./model/*.proto  * означает, что все файлы с расширением proto находятся в текущей директории

# Команда для генерации protobuf сервисов:
    protoc --go-grpc_out=../pb ./service/*.proto  

### 1. protoc - команда для запуска из env_path
### 2. --go-grpc_out= path, куда будет генерироваться {name}.pb.go
#### так как в proto-файлах указан go_package = "/service", то конечный путь будет: ..pb/service/
### 3. ./service/*.proto  * означает, что все файлы с расширением proto находятся в текущей директории