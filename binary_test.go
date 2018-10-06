package tinfer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryUnify(t *testing.T) {
	tt, ttt := Type(NewBinary("")), Type(NewBinary(""))
	assert.Nil(t, tt.Unify(ttt))
}

func TestBinaryUnifyError(t *testing.T) {
	for _, ts := range [][2]Type{
		{NewBinary(""), NewNumber("")},
		{NewNumber(""), NewBinary("")},
	} {
		assert.Error(t, ts[0].Unify(ts[1]))
	}
}
