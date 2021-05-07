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

### Steps to use the application
As the application features gRPC gateway, the services can be accessible by both using and rpc client and also using an REST API client.
 - Using curl
  ``` shell
   curl -X POST -k http://localhost:8081/v1/user/fetchByID -d '{"id":2}'
  ```
or 
  ``` shell
   curl -X POST -k http://localhost:8081/v1/user/fetch -d '{"id":[1,2]}'
  ```
 - You can also use a rpc client to use the service. You can use the proto files to create a client or any cli can be used to connect with the server
 - You can also use cli like [evans](https://github.com/ktr0731/evans) in order to use the rpc server.