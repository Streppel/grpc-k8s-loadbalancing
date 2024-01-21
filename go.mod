module github.com/Streppel/grpc-k8s-loadbalancing

go 1.21

require google.golang.org/grpc v1.60.1

replace github.com/Streppel/grpc-k8s-loadbalancing => ./proto/github.com/Streppel/grpc-k8s-loadbalancing

require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.16.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231002182017-d307bd883b97 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
)

