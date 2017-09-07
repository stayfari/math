package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoundDown(t *testing.T) {
	raw := 0.49
	result := Round(raw)
	assert.Equal(t, result, 0.0, "Round Down failed")
}

func TestRoundUp(t *testing.T) {
	raw := 0.50
	result := Round(raw)
	assert.Equal(t, result, 1.0, "Round Down failed")
}

func TestRoundToPlaces(t *testing.T) {
	raw := 0.123456
	result := RoundToPlaces(raw, 5)
	assert.Equal(t, result, 0.12346, "Round to 5 places failed")
	result = RoundToPlaces(raw, 4)
	assert.Equal(t, result, 0.1235, "Round to 4 places failed")
	result = RoundToPlaces(raw, 3)
	assert.Equal(t, result, 0.123, "Round to 3 places failed")
	result = RoundToPlaces(raw, 2)
	assert.Equal(t, result, 0.12, "Round to 2 places failed")
}
