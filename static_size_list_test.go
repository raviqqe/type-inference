package tinfer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStaticSizeListUnify(t *testing.T) {
	for _, ts := range [][2]Type{
		{
			NewStaticSizeList(nil, ""),
			NewStaticSizeList(nil, ""),
		},
		{
			NewStaticSizeList(nil, ""),
			NewVariable(""),
		},
		{
			NewStaticSizeList([]Type{NewNumber("")}, ""),
			NewStaticSizeList([]Type{NewNumber("")}, ""),
		},
		{
			NewStaticSizeList([]Type{NewNumber("")}, ""),
			NewStaticSizeList([]Type{NewVariable("")}, ""),
		},
		{
			NewStaticSizeList([]Type{NewNumber(""), NewNumber("")}, ""),
			NewStaticSizeList([]Type{NewNumber(""), NewVariable("")}, ""),
		},
		{
			NewStaticSizeList([]Type{NewNumber(""), NewNumber("")}, ""),
			NewDynamicSizeList(NewNumber(""), ""),
		},
	} {
		assert.Nil(t, ts[0].Unify(ts[1]))
	}
}

func TestStaticSizeListIsList(t *testing.T) {
	NewStaticSizeList([]Type{NewNumber("")}, "").isList()
}
