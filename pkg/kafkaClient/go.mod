module github.com/onosproject/analytics/pkg/kafkaClient

go 1.17

require github.com/segmentio/kafka-go v0.4.29

require (
	github.com/klauspost/compress v1.14.2 // indirect
	github.com/onosproject/analytics/pkg/logger v0.0.0-unpublished
	github.com/pierrec/lz4/v4 v4.1.14 // indirect

)

replace github.com/onosproject/analytics/pkg/logger v0.0.0-unpublished => ../logger
