
Generate protoc files
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