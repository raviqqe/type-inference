package tinfer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNullAccept(t *testing.T) {
	tt, ttt := Type(NewNull("")), Type(NewNull(""))
	assert.Nil(t, tt.Accept(ttt))
}

func TestNullAcceptError(t *testing.T) {
	for _, ts := range [][2]Type{
		{NewNull(""), NewNumber("")},
		{NewNumber(""), NewNull("")},
	} {
		assert.Error(t, ts[0].Accept(ts[1]))
	}
}

func TestNullCanAccept(t *testing.T) {
	assert.True(t, NewNull("").CanAccept(NewNull("")))
	assert.False(t, NewNull("").CanAccept(NewNumber("")))
}
