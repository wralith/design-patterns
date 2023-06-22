package singleton

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewLogger(t *testing.T) {
	instance1 := NewLogger()
	instance2 := NewLogger()

	assert.Equal(t, instance1, instance2)
	assert.Equal(t, &instance1, &instance2)
}
