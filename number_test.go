package tinfer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberAccept(t *testing.T) {
	tt, ttt := Type(NewNumber("")), Type(NewNumber(""))
	assert.Nil(t, tt.Accept(ttt))
}

func TestNumberAcceptError(t *testing.T) {
	tt, ttt := Type(NewNumber("")), Type(NewFunction(nil, NewNumber(""), ""))
	assert.Error(t, tt.Accept(ttt))
}

func TestNumberCanAccept(t *testing.T) {
	assert.True(t, NewNumber("").CanAccept(NewNumber("")))
	assert.False(t, NewNumber("").CanAccept(NewString("")))
}
