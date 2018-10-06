package tinfer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringAccept(t *testing.T) {
	tt, ttt := Type(NewString("")), Type(NewString(""))
	assert.Nil(t, tt.Accept(ttt))
}

func TestStringAcceptError(t *testing.T) {
	for _, ts := range [][2]Type{
		{NewString(""), NewNumber("")},
		{NewNumber(""), NewString("")},
	} {
		assert.Error(t, ts[0].Accept(ts[1]))
	}
}
