module github.com/OpenStars/BackendService

go 1.15

require (
	github.com/OpenStars/EtcdBackendService v0.0.0-20200902010328-9a9cc534fe8e
	github.com/OpenStars/GoEndpointManager v0.0.0-20200513065934-c2f3d8399632
	github.com/OpenStars/thriftpoolv2 v0.0.0-20200306081147-89225e956ca9 // indirect
	github.com/apache/thrift v0.13.0
	github.com/coreos/etcd v3.3.25+incompatible // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd v0.0.0-00010101000000-000000000000 // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/elastic/go-elasticsearch v0.0.0
	github.com/go-telegram-bot-api/telegram-bot-api v4.6.4+incompatible
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/google/uuid v1.1.2 // indirect
	github.com/technoweenie/multipartstreamer v1.0.1 // indirect
	go.etcd.io/etcd v3.3.25+incompatible
	go.uber.org/zap v1.16.0 // indirect
	google.golang.org/grpc v1.26.0 // indirect

)

replace github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0
