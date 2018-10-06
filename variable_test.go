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
		assert.Nil(t, ts[0].Unify(ts[1]))
	}
}

func TestVariableUnifyInferredType(t *testing.T) {
	tt, ttt := Type(NewVariable("")), Type(NewDynamicSizeList(NewVariable(""), ""))

	assert.Nil(t, tt.Unify(ttt))
	assert.Equal(t, NewDynamicSizeList(NewVariable(""), ""), ttt)
	assert.Nil(t, tt.Unify(NewDynamicSizeList(NewNumber(""), "")))
	assert.Equal(t, NewDynamicSizeList(&Variable{NewNumber(""), ""}, ""), ttt)
}

func TestVariableLocation(t *testing.T) {
	assert.Equal(t, "foo", NewVariable("foo").Location())
}
