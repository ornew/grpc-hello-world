# gRPC Hello World for debug

It can be used to check the communication of gRPC.

```
go run github.com/ornew/cmd/client -address localhost:8081 -name john
```

If name is `arata`, return InvalidArgument. If name is `furukawa`, return other message.
