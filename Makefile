IMAGE ?= ornew/grpc-hello-world

gen:
	go generate ./...

docker-build:
	docker build -t $(IMAGE) .

docker-push:
	docker push $(IMAGE)
