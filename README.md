# gRPC Hello World for debug

It can be used to check the communication of gRPC.

```
docker run --rm -it -p 8081:8081 ornew/grpc-hello-world
go run github.com/ornew/grpc-hello-world/cmd/client -address localhost:8081 -name john
```

If name is `arata`, return InvalidArgument. If name is `furukawa`, return other message.
