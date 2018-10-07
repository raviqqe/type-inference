package tinfer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUnion(t *testing.T) {
	for _, ts := range [][]Type{
		{NewNumber(""), NewNumber("")},
		{NewNumber(""), NewNumber(""), NewNumber("")},
	} {
		assert.NotPanics(t, func() { NewUnion(ts, "") })
	}
}

func TestNewUnionPanic(t *testing.T) {
	assert.Panics(t, func() { NewUnion(nil, "") })
	assert.Panics(t, func() { NewUnion([]Type{NewNumber("")}, "") })
}

func TestUnionAccept(t *testing.T) {
	for _, ts := range [][2]Type{
		{
			NewUnion([]Type{NewNumber(""), NewString("")}, ""),
			NewNumber(""),
		},
		{
			NewUnion([]Type{NewNumber(""), NewString("")}, ""),
			NewUnion([]Type{NewNumber(""), NewString("")}, ""),
		},
		{
			NewUnion([]Type{NewNumber(""), NewString(""), NewNull("")}, ""),
			NewUnion([]Type{NewNumber(""), NewString("")}, ""),
		},
		{
			NewMapObject(NewNumber(""), ""),
			NewUnion(
				[]Type{
					NewRecordObject(map[string]Type{"foo": NewNumber("")}, ""),
					NewRecordObject(map[string]Type{"bar": NewNumber("")}, ""),
				},
				"",
			),
		},
		{
			NewDynamicSizeList(NewNumber(""), ""),
			NewUnion(
				[]Type{
					NewStaticSizeList([]Type{NewNumber("")}, ""),
					NewStaticSizeList([]Type{NewNumber(""), NewNumber("")}, ""),
				},
				"",
			),
		},
		// TODO: Check all combinations of member types.
		// {
		// 	NewUnion([]Type{NewNumber(""), NewVariable("")}, ""),
		// 	NewUnion([]Type{NewNumber(""), NewString("")}, ""),
		// },
	} {
		assert.Nil(t, ts[0].Accept(ts[1]))
	}
}

func TestUnionCanAcceptPanic(t *testing.T) {
	assert.Panics(t, func() {
		NewUnion([]Type{NewNumber(""), NewString("")}, "").CanAccept(NewNumber(""))
	})
}

func TestUnionAcceptError(t *testing.T) {
	for _, ts := range [][2]Type{
		{
			NewUnion([]Type{NewNumber(""), NewString("")}, ""),
			NewNull(""),
		},
		{
			NewUnion([]Type{NewNumber(""), NewString("")}, ""),
			NewUnion([]Type{NewNumber(""), NewString(""), NewNull("")}, ""),
		},
		{
			NewMapObject(NewNumber(""), ""),
			NewUnion(
				[]Type{
					NewRecordObject(map[string]Type{"foo": NewNumber("")}, ""),
					NewRecordObject(map[string]Type{"bar": NewString("")}, ""),
				},
				"",
			),
		},
		{
			NewDynamicSizeList(NewNumber(""), ""),
			NewUnion(
				[]Type{
					NewStaticSizeList([]Type{NewNumber("")}, ""),
					NewStaticSizeList([]Type{NewNumber(""), NewString("")}, ""),
				},
				"",
			),
		},
	} {
		assert.NotNil(t, ts[0].Accept(ts[1]))
	}
}
