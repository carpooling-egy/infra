package internal

import (
	"context"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/rs/zerolog/log"
)

type JetStreamManager struct {
	js             jetstream.JetStream
	streamName     string
	subjectPattern string
	consumerName   string
	consumerSubj   string
}

func NewJetStreamManager(js jetstream.JetStream, streamName, subjectPattern, consumerName, consumerSubj string) *JetStreamManager {
	return &JetStreamManager{
		js:             js,
		streamName:     streamName,
		subjectPattern: subjectPattern,
		consumerName:   consumerName,
		consumerSubj:   consumerSubj,
	}
}

func (m *JetStreamManager) Setup(ctx context.Context) error {
	stream, err := m.js.CreateOrUpdateStream(ctx, jetstream.StreamConfig{
		Name:     m.streamName,
		Subjects: []string{m.subjectPattern},
	})
	if err != nil {
		return err
	}
	log.Info().Str("stream", m.streamName).Msg("Stream created or updated")

	_, err = stream.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
		Name:          m.consumerName,
		Durable:       m.consumerName,
		AckPolicy:     jetstream.AckExplicitPolicy,
		FilterSubject: m.consumerSubj,
	})
	if err != nil {
		return err
	}
	log.Info().Str("consumer", m.consumerName).Msg("Consumer created or updated")
	return nil
}
