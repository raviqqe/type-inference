package tinfer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVariableAccept(t *testing.T) {
	for _, ts := range [][2]Type{
		{NewVariable(""), NewVariable("")},
		{NewVariable(""), NewNumber("")},
	} {
		assert.Nil(t, ts[0].Accept(ts[1]))
	}
}

func TestVariableAcceptInferredType(t *testing.T) {
	tt, ttt := Type(NewVariable("")), Type(NewDynamicSizeList(NewVariable(""), ""))

	assert.Nil(t, tt.Accept(ttt))
	assert.Equal(t, NewDynamicSizeList(NewVariable(""), ""), ttt)
	assert.Nil(t, tt.Accept(NewDynamicSizeList(NewNumber(""), "")))
	assert.Equal(t, NewDynamicSizeList(&Variable{NewNumber(""), ""}, ""), ttt)
}

func TestVariableLocation(t *testing.T) {
	assert.Equal(t, "foo", NewVariable("foo").Location())
}
