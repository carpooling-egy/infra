package main

import (
	"InitJetStream/internal"
	"context"
	"os"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Logger()

	cfg := internal.LoadConfig()

	var opts []nats.Option
	if cfg.NATSUser != "" && cfg.NATSPass != "" {
		opts = append(opts, nats.UserInfo(cfg.NATSUser, cfg.NATSPass))
	}

	nc, err := nats.Connect(cfg.NATSURL, opts...)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to NATS")
	}
	defer func(nc *nats.Conn) {
		err := nc.Drain()
		if err != nil {
			log.Error().Err(err).Msg("Failed to drain connection")
		}
	}(nc)

	js, err := jetstream.New(nc)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create JetStream context")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	manager := internal.NewJetStreamManager(js, cfg.StreamName, cfg.SubjectPattern, cfg.ConsumerName, cfg.ConsumerSubject)
	if err := manager.Setup(ctx); err != nil {
		log.Fatal().Err(err).Msg("JetStream setup failed")
	}
}
