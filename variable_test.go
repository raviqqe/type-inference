package tinfer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVariableUnify(t *testing.T) {
	for _, ts := range [][2]Type{
		{NewVariable(""), NewVariable("")},
		{NewVariable(""), NewNumber("")},
		{NewNumber(""), NewVariable("")},
	} {
		err := ts[0].Unify(&ts[0], &ts[1])

		assert.Nil(t, err)
		assert.Equal(t, ts[0], ts[1])
	}
}

func TestVariableLocation(t *testing.T) {
	assert.Equal(t, "foo", NewVariable("foo").Location())
}
