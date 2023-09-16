package logger

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSourceLocation(t *testing.T) {
	t.Parallel()

	got := SourceLocation(runtime.Caller(0)).Interface.(*source)

	assert.Contains(t, got.File, "go-logger/source_test.go")
	assert.Equal(t, "13", got.Line)
	assert.Contains(t, got.Function, "logger.TestSourceLocation")
}

func TestNewSource(t *testing.T) {
	t.Parallel()

	got := newSource(runtime.Caller(0))

	assert.Contains(t, got.File, "go-logger/source_test.go")
	assert.Equal(t, "23", got.Line)
	assert.Contains(t, got.Function, "logger.TestNewSource")
}
