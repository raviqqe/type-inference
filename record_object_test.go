package tinfer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecordObjectUnify(t *testing.T) {
	for _, ts := range [][2]Type{
		{
			NewRecordObject(map[string]Type{"foo": NewNumber("")}, ""),
			NewRecordObject(map[string]Type{"foo": NewNumber("")}, ""),
		},
		{
			NewRecordObject(map[string]Type{"foo": NewNumber("")}, ""),
			NewRecordObject(map[string]Type{"foo": NewNumber(""), "bar": NewString("")}, ""),
		},
		{
			NewRecordObject(map[string]Type{"foo": NewNumber(""), "bar": NewString("")}, ""),
			NewRecordObject(map[string]Type{"foo": NewNumber("")}, ""),
		},
		{
			NewRecordObject(map[string]Type{"foo": NewNumber("")}, ""),
			NewMapObject(NewNumber(""), ""),
		},
		{
			NewRecordObject(map[string]Type{"foo": NewNumber("")}, ""),
			NewVariable(""),
		},
	} {
		assert.Nil(t, ts[0].Unify(&ts[0], &ts[1]))
		assert.Equal(t, ts[0], ts[1])
	}
}

func TestRecordObjectUnifyError(t *testing.T) {
	for _, ts := range [][2]Type{
		{
			NewRecordObject(map[string]Type{"foo": NewNumber("")}, ""),
			NewRecordObject(map[string]Type{"foo": NewString("")}, ""),
		},
		{
			NewNumber(""),
			NewRecordObject(map[string]Type{"foo": NewString("")}, ""),
		},
	} {
		assert.Error(t, ts[0].Unify(&ts[0], &ts[1]))
	}
}

func TestRecordObjectContain(t *testing.T) {
	for _, os := range [][2]*RecordObject{
		{
			NewRecordObject(nil, ""),
			NewRecordObject(nil, ""),
		},
		{
			NewRecordObject(map[string]Type{"foo": NewNumber("")}, ""),
			NewRecordObject(map[string]Type{"foo": NewNumber("")}, ""),
		},
		{
			NewRecordObject(map[string]Type{"foo": NewNumber(""), "bar": NewNumber("")}, ""),
			NewRecordObject(map[string]Type{"foo": NewNumber(""), "bar": NewNumber("")}, ""),
		},
	} {
		assert.True(t, os[0].contain(os[1]))
	}

	for _, os := range [][2]*RecordObject{
		{
			NewRecordObject(nil, ""),
			NewRecordObject(map[string]Type{"foo": NewNumber("")}, ""),
		},
		{
			NewRecordObject(map[string]Type{"foo": NewNumber("")}, ""),
			NewRecordObject(map[string]Type{"foo": NewNumber(""), "bar": NewNumber("")}, ""),
		},
	} {
		assert.True(t, os[0].contain(os[1]))
		assert.False(t, os[1].contain(os[0]))
	}
}

func TestRecordObjectIsObject(t *testing.T) {
	NewRecordObject(nil, "").isObject()
}
