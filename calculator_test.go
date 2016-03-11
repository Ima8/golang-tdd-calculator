package calculator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAplusB(t *testing.T) {
	assert.Equal(t, plus(1, 2), 3)
	assert.Equal(t, plus(2, 2), 4)
}

func TestAdelB(t *testing.T) {
	assert.Equal(t, del(1, 1), 0)
	assert.Equal(t, del(2, 3), -1)
	assert.Equal(t, del(1, 0), 1)
}
func TestAMultiB(t *testing.T) {
	assert.Equal(t, multi(1, 1), 1)
	assert.Equal(t, multi(2, 3), 6)
	assert.Equal(t, multi(1, 0), 0)
}
func TestAdivB(t *testing.T) {
	assert.Equal(t, div(1, 1), 1)
	assert.Equal(t, div(2, 2), 1)
	assert.Equal(t, div(3, 2), 1)
}
