module github.com/onosproject/analytics

go 1.17

require (
	github.com/onosproject/analytics/pkg/kafkaClient v0.0.0-unpublished
	github.com/onosproject/analytics/pkg/logger v0.0.0-unpublished
	github.com/onosproject/analytics/pkg/messages v0.0.0-20220307215111-c7d1cd474463
)

require (
	github.com/klauspost/compress v1.14.2 // indirect
	github.com/pierrec/lz4/v4 v4.1.14 // indirect
	github.com/segmentio/kafka-go v0.4.29 // indirect
)

replace github.com/onosproject/analytics/pkg/kafkaClient v0.0.0-unpublished => ./pkg/kafkaClient

replace github.com/onosproject/analytics/pkg/logger v0.0.0-unpublished => ./pkg/logger
