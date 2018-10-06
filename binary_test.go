package tinfer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryAccept(t *testing.T) {
	tt, ttt := Type(NewBinary("")), Type(NewBinary(""))
	assert.Nil(t, tt.Accept(ttt))
}

func TestBinaryAcceptError(t *testing.T) {
	for _, ts := range [][2]Type{
		{NewBinary(""), NewNumber("")},
		{NewNumber(""), NewBinary("")},
	} {
		assert.Error(t, ts[0].Accept(ts[1]))
	}
}
