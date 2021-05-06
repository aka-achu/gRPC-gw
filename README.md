## gRPC-gw
gRPC-gw is a simple go gRPC server with support for REST API

### Steps to run the application
- Build from source
  ```shell
  # download the dependencies
  go mod dowload
  # build the binary
  go build main.go -o app
  # execution permission to the binary
  chmod +x ./app
  # run the application
  ./app
  ```
- Using Docker
  ```shell
  # build the docker image
  docker build -t app:latest -f Dockerfile .
  # run the container
  docker run --rm -d -p 8080:8080 -p 8081:8081 --name app app:latest
  ```