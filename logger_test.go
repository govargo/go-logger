package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProductionConfig(t *testing.T) {
	npc := NewProductionConfig()
	assert.Equal(t, "info", npc.Level.String())
	assert.Equal(t, false, npc.Development)
	assert.Equal(t, "json", npc.Encoding)
}

func TestDebug(t *testing.T) {
	Debug("test debug")
}

func TestInfo(t *testing.T) {
	Info("test info")
}

func TestWarn(t *testing.T) {
	Warn("test warning")
}

func TestError(t *testing.T) {
	Error("test error")
}

func TestInterceptorLogger(t *testing.T) {
	l := InterceptorLogger(logger)
	assert.NotEmpty(t, l)
}

func TestGetGRPCLogger(t *testing.T) {
	gl := GetGRPCLogger()
	assert.NotEmpty(t, gl)
}
