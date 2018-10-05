package tinfer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberUnify(t *testing.T) {
	tt, ttt := Type(NewNumber("")), Type(NewNumber(""))
	assert.Nil(t, tt.Unify(&tt, &ttt))
}

func TestNumberUnifyError(t *testing.T) {
	tt, ttt := Type(NewNumber("")), Type(NewFunction(nil, NewNumber(""), ""))
	assert.Error(t, tt.Unify(&tt, &ttt))
}
