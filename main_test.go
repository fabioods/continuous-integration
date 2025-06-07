package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSoma(t *testing.T) {
	result := Soma(3, 4)
	assert.Equal(t, 7, result, "Expected sum of 3 and 4 to be 7")
}
