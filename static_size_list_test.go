package tinfer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStaticSizeListAccept(t *testing.T) {
	for _, ts := range [][2]Type{
		{
			NewStaticSizeList(nil, ""),
			NewStaticSizeList(nil, ""),
		},
		{
			NewStaticSizeList([]Type{NewNumber("")}, ""),
			NewStaticSizeList([]Type{NewNumber("")}, ""),
		},
		{
			NewStaticSizeList([]Type{NewVariable("")}, ""),
			NewStaticSizeList([]Type{NewNumber("")}, ""),
		},
		{
			NewStaticSizeList([]Type{NewNumber(""), NewVariable("")}, ""),
			NewStaticSizeList([]Type{NewNumber(""), NewNumber("")}, ""),
		},
		{
			NewStaticSizeList([]Type{NewNumber(""), NewNumber("")}, ""),
			NewDynamicSizeList(NewNumber(""), ""),
		},
	} {
		assert.Nil(t, ts[0].Accept(ts[1]))
	}
}

func TestStaticSizeListAcceptError(t *testing.T) {
	for _, ts := range [][2]Type{
		{
			NewStaticSizeList(nil, ""),
			NewNumber(""),
		},
		{
			NewNumber(""),
			NewStaticSizeList(nil, ""),
		},
		{
			NewStaticSizeList([]Type{NewNumber("")}, ""),
			NewStaticSizeList([]Type{NewNumber(""), NewNumber("")}, ""),
		},
		{
			NewStaticSizeList([]Type{NewNumber("")}, ""),
			NewStaticSizeList([]Type{NewString("")}, ""),
		},
		{
			NewStaticSizeList([]Type{NewNumber("")}, ""),
			NewDynamicSizeList(NewString(""), ""),
		},
	} {
		assert.Error(t, ts[0].Accept(ts[1]))
	}
}

func TestStaticSizeListIsList(t *testing.T) {
	NewStaticSizeList([]Type{NewNumber("")}, "").isList()
}
