package tinfer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInferenceErrorError(t *testing.T) {
	assert.Equal(t,
		newInferenceError("type error", "foo.go:42:2049").Error(),
		"foo.go:42:2049: type error")
}
