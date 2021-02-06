module github.com/OpenStars/BackendService

go 1.15

replace go.etcd.io/etcd/v3 => /home/lehaisonmath6/go/src/go.etcd.io/etcd

require (
	github.com/OpenStars/EtcdBackendService v0.0.0-20201021070238-0d4c60000fbe
	github.com/apache/thrift v0.13.0
	github.com/bluele/gcache v0.0.0-20190518031135-bc40bd653833
	github.com/elastic/go-elasticsearch/v8 v8.0.0-20201104130540-2e1f801663c6
	github.com/go-telegram-bot-api/telegram-bot-api v4.6.4+incompatible
	github.com/rtt/Go-Solr v0.0.0-20190512221613-64fac99dcae2
	github.com/segmentio/kafka-go v0.4.9
	github.com/technoweenie/multipartstreamer v1.0.1 // indirect
)
