package contexts_test

import (
	"context"
	"testing"

	"github.com/ramamimu/go-everything/3-concurrency/contexts"
	"github.com/stretchr/testify/assert"
)

func TestFetchWithContext(t *testing.T) {
	msg, err := contexts.FetchWithContext(context.Background(), "hello world")
	assert.Containsf(t, msg, "succeed-hello world", "expected 'succeed-hello world' but got '%s'", msg)
	assert.NoError(t, err, "expected no error")
}
