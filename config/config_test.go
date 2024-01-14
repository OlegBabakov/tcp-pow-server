package config

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewServerConfig(t *testing.T) {
	c, err := NewConfig(context.Background(), ServerConfig{})
	require.NoError(t, err)
	assert.NotEmpty(t, c.Addr)
	assert.NotEmpty(t, c.KeepAlive)
	assert.NotEmpty(t, c.ConnDeadline)
}

func TestNewClientConfig(t *testing.T) {
	c, err := NewConfig(context.Background(), ClientConfig{})
	require.NoError(t, err)
	assert.NotEmpty(t, c.ServerAddr)
	assert.NotEmpty(t, c.KeepAlive)
	assert.NotEmpty(t, c.RequestCount)
}
