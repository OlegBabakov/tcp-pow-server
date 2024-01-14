package config

import (
	"context"
	"time"

	"github.com/sethvargo/go-envconfig"
)

// ServerConfig server configuration.
type ServerConfig struct {
	Addr              string        `env:"SERVER_ADDR,default=0.0.0.0:8000"`
	KeepAlive         time.Duration `env:"SERVER_KEEP_ALIVE,default=15s"`
	ConnDeadline      time.Duration `env:"SERVER_CONN_DEADLINE,default=10s"`
	ConnAcceptTimeout time.Duration `env:"SERVER_ACCEPT_TIMEOUT,default=3s"`
	Workers           int           `env:"SERVER_WORKERS,default=1"`
	DefaultConfig
}

// ClientConfig client configuration.
type ClientConfig struct {
	ServerAddr   string        `env:"SERVER_ADDR,default=127.0.0.0:8000"`
	RequestCount int           `env:"CLIENT_REQUEST_COUNT,default=100"`
	KeepAlive    time.Duration `env:"CLIENT_KEEP_ALIVE,default=15s"`
	DefaultConfig
}

// DefaultConfig for client and server.
type DefaultConfig struct {
	Logger LoggerConfig `env:",prefix=LOGGER_"`
	Pow    PowConfig    `env:",prefix=POW_"`
}

// LoggerConfig for logger configuration.
type LoggerConfig struct {
	DisableCaller     bool   `env:"CALLER,default=false"`
	DisableStacktrace bool   `env:"STACKTRACE,default=false"`
	Level             string `env:"LEVEL,default=debug"`
}

// PowConfig for PoW configuration.
type PowConfig struct {
	Complexity uint64 `env:"COMPLEXITY,default=20"`
}

// NewConfig generic for creates a new client or server config.
func NewConfig[C any](ctx context.Context, config C) (*C, error) {
	if err := envconfig.Process(ctx, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
