package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRedis(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("Panic error:", r)
		}
	}()

	_ = GetRedis()

	assert.True(t, true)
}
