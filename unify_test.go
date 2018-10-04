package tinfer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnify(t *testing.T) {
	tt, ttt := Type(NewNumber("")), Type(NewNumber(""))
	assert.Nil(t, Unify(&tt, &ttt))
}

func TestUnifyError(t *testing.T) {
	tt, ttt := Type(NewNumber("")), Type(NewFunction(nil, NewNumber(""), ""))
	assert.NotNil(t, Unify(&tt, &ttt))
}
