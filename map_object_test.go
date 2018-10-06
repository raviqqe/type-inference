package tinfer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestObjectUnify(t *testing.T) {
	for _, ts := range [][2]Type{
		{
			NewMapObject(NewNumber(""), ""),
			NewMapObject(NewNumber(""), ""),
		},
		{
			NewMapObject(NewNumber(""), ""),
			NewVariable(""),
		},
		{
			NewMapObject(NewNumber(""), ""),
			NewMapObject(NewVariable(""), ""),
		},
		{
			NewMapObject(NewNumber(""), ""),
			NewRecordObject(map[string]Type{"foo": NewNumber(""), "bar": NewNumber("")}, ""),
		},
	} {
		assert.Nil(t, ts[0].Unify(ts[1]))
	}
}

func TestObjectUnifyError(t *testing.T) {
	for _, ts := range [][2]Type{
		{
			NewNumber(""),
			NewMapObject(NewString(""), ""),
		},
		{
			NewMapObject(NewNumber(""), ""),
			NewMapObject(NewString(""), ""),
		},
		{
			NewMapObject(NewNumber(""), ""),
			NewRecordObject(map[string]Type{"foo": NewNumber(""), "bar": NewString("")}, ""),
		},
	} {
		assert.Error(t, ts[0].Unify(ts[1]))
	}
}

func TestMapObjectIsObject(t *testing.T) {
	NewMapObject(NewNumber(""), "").isObject()
}
