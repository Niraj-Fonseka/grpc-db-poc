guide : https://medium.com/pantomath/how-we-use-grpc-to-build-a-client-server-system-in-go-dd20045fa1c2


```
protoc -I api/ -I${GOPATH}/src --go_out=plugins=grpc:api api/api.proto
```

Build Server 
```
go build -i -v -o bin/server $GOPATH/src/GRPCThings/GoClientServer/server
```


Build Client 
```
go build -i -v -o bin/client $GOPATH/src/GRPCThings/GoClientServer/client
```