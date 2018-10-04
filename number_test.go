package tinfer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberUnify(t *testing.T) {
	for _, ts := range [][2]Type{{NewNumber(""), NewNumber("")}} {
		assert.Nil(t, ts[0].Unify(&ts[0], &ts[1]))
	}
}
