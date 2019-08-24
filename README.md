### golang_grpc_demo

###### Create Client Server Skeleton using the proto compiler
```$xslt
protoc grpcpb/app.proto --go_out=plugins=grpc:.
```
###### Create Docker Image
```$xslt
docker build -t mygrpcservice:latest . -f Dockerfile 
```

###### Run Docker Container
```$xslt
docker run -it --name mygrpcservice -p 50051:50051 mygrpcservice:latest

GRPC Server starting!!
[GetServerInfo] serving GetServerInfo request
```

###### Test using grpc client
```$xslt
go run grpc_client/client.go 
[callGetServerInfo]  client is about to send GetServerInfo
[callGetServerInfo] Response : hostname : a350fb0803ac current server time : 2019-08-24 05:10:00.239490482 +0000 UTC m=+9.922789757
```