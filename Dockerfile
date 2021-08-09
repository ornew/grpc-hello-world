ARG GO_PACKAGE=github.com/ornew/grpc-hello-world

FROM golang:1.16 as build

ARG GO_PACKAGE

WORKDIR /go/src/${GO_PACKAGE}

COPY . .

RUN go install ./cmd/client && go install ./cmd/server

FROM gcr.io/distroless/base

COPY --from=build /go/bin/client /client
COPY --from=build /go/bin/server /server

CMD ["/server"]
