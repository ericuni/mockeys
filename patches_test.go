package mockeys

import (
	"testing"

	"github.com/bytedance/mockey"
	"github.com/stretchr/testify/assert"
)

func Foo(x, y int) int {
	return x + y
}

func Bar(x, y int) int {
	return x * y
}

func TestApply(t *testing.T) {
	assert := assert.New(t)

	patches := NewPatches()
	defer patches.Reset()

	// to
	mock := mockey.Mock(Foo).To(func(x, y int) int {
		return x - y
	})
	patches.Apply(mock)

	// return
	mock = mockey.Mock(Bar).Return(10)
	patches.Apply(mock)

	assert.Equal(-1, Foo(1, 2))
	assert.Equal(10, Bar(1, 2))
}

var batchSize int = 10

func TestApplyVar(t *testing.T) {
	assert := assert.New(t)

	patches := NewPatches()
	defer patches.Reset()

	mock := mockey.MockValue(&batchSize).To(20)
	patches.ApplyVar(mock)

	assert.Equal(20, batchSize)
}
