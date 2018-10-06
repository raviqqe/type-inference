package tinfer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNullUnify(t *testing.T) {
	tt, ttt := Type(NewNull("")), Type(NewNull(""))
	assert.Nil(t, tt.Unify(ttt))
}

func TestNullUnifyError(t *testing.T) {
	for _, ts := range [][2]Type{
		{NewNull(""), NewNumber("")},
		{NewNumber(""), NewNull("")},
	} {
		assert.Error(t, ts[0].Unify(ts[1]))
	}
}
