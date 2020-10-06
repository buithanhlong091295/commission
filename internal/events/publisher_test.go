package events

import (
	"testing"

	"github.com/richard-xtek/go-grpc-micro-kit/kafka"

	"github.com/richard-xtek/go-grpc-micro-kit/log"
)

func loadTestingPublisher(t *testing.T) *Publisher {
	logger := log.NewDevelopFactory("publisher_test")
	publishCfg := kafka.PublisherConfig{
		Brokers:   []string{"localhost:9092"},
		Marshaler: kafka.DefaultMarshaler{},
	}

	publisher, err := kafka.NewPublisher(publishCfg, logger)
	if err != nil {
		t.Fatal(err)
	}

	return &Publisher{pub: publisher}
}
