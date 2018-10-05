package tinfer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringUnify(t *testing.T) {
	tt, ttt := Type(NewString("")), Type(NewString(""))
	assert.Nil(t, tt.Unify(&tt, &ttt))
}

func TestStringUnifyError(t *testing.T) {
	for _, ts := range [][2]Type{
		{NewString(""), NewNumber("")},
		{NewNumber(""), NewString("")},
	} {
		assert.Error(t, ts[0].Unify(&ts[0], &ts[1]))
	}
}
