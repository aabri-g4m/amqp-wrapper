package plugin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigInitDefaults(t *testing.T) {
	cfg := GlobalCfg{}

	cfg.InitDefaults()

	assert.Equal(t, "tcp://localhost:8088", cfg.Addr)
}
