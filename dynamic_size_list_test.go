package tinfer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDynamicSizeListUnify(t *testing.T) {
	for _, ts := range [][2]Type{
		{
			NewDynamicSizeList(NewNumber(""), ""),
			NewDynamicSizeList(NewNumber(""), ""),
		},
		{
			NewDynamicSizeList(NewNumber(""), ""),
			NewDynamicSizeList(NewVariable(""), ""),
		},
		{
			NewDynamicSizeList(NewDynamicSizeList(NewVariable(""), ""), ""),
			NewDynamicSizeList(NewDynamicSizeList(NewNumber(""), ""), ""),
		},
		{
			NewDynamicSizeList(NewNumber(""), ""),
			NewStaticSizeList([]Type{NewNumber("")}, ""),
		},
		{
			NewDynamicSizeList(NewNumber(""), ""),
			NewStaticSizeList([]Type{NewNumber(""), NewVariable("")}, ""),
		},
	} {
		assert.Nil(t, ts[0].Unify(&ts[0], &ts[1]))
		assert.Equal(t, ts[0], ts[1])
	}
}

func TestDynamicSizeListUnifyError(t *testing.T) {
	for _, ts := range [][2]Type{
		{
			NewDynamicSizeList(NewNumber(""), ""),
			NewNumber(""),
		},
		{
			NewDynamicSizeList(NewNumber(""), ""),
			NewStaticSizeList([]Type{NewNumber(""), NewDynamicSizeList(NewNumber(""), "")}, ""),
		},
	} {
		assert.Error(t, ts[0].Unify(&ts[0], &ts[1]))
	}
}

func TestDynamicSizeListIsList(t *testing.T) {
	NewDynamicSizeList(NewNumber(""), "").isList()
}
