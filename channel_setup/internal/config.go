package internal

import (
	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog/log"
	"os"
)

type Config struct {
	NATSURL         string
	NATSUser        string
	NATSPass        string
	StreamName      string
	SubjectPattern  string
	ConsumerName    string
	ConsumerSubject string
}

func LoadConfig() Config {
	user := os.Getenv("NATS_USER")
	pass := os.Getenv("NATS_PASSWORD")
	if user == "" || pass == "" {
		log.Warn().Str("NATS_USER", user).Str("NATS_PASS", mask(pass)).Msg("NATS credentials not fully set; proceeding without authentication")
	}

	return Config{
		NATSURL:         getEnv("NATS_URL", nats.DefaultURL),
		NATSUser:        user,
		NATSPass:        pass,
		StreamName:      getEnv("NATS_STREAM_NAME", "MATCHED_REQUESTS"),
		SubjectPattern:  getEnv("NATS_STREAM_SUBJECT", "matched_requests.*"),
		ConsumerName:    getEnv("NATS_CONSUMER_NAME", "MATCHED_REQUESTS_PROCESSOR"),
		ConsumerSubject: getEnv("NATS_CONSUMER_SUBJECT", "matched_requests.results"),
	}
}
