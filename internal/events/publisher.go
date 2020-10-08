package events

import (
	"github.com/richard-xtek/go-grpc-micro-kit/kafka"
)

// NewPublisher ...
func NewPublisher(pub *kafka.Publisher) *Publisher {
	return &Publisher{pub}
}

// Publisher ...
type Publisher struct {
	pub *kafka.Publisher
}
