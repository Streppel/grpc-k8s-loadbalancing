.PHONY: setup-kind create-cluster deploy-server deploy-client build-images deploy

create-cluster: setup-kind
	kind create cluster --name grpc-demo

deploy-server:
	kubectl apply -f grpc-server.yaml

deploy-client:
	kubectl apply -f grpc-client.yaml

build-images:
	docker build -t grpc-server:local -f server/Dockerfile .
	docker build -t grpc-client:local -f client/Dockerfile .

load-images: build-images
	kind load docker-image grpc-server:local --name grpc-demo
	kind load docker-image grpc-client:local --name grpc-demo

deploy: deploy-server deploy-client

.PHONY: clean
clean:
	kind delete cluster --name grpc-demo
