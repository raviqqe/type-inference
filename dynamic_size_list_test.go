package tinfer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDynamicSizeListAccept(t *testing.T) {
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
		assert.Nil(t, ts[0].Accept(ts[1]))
	}
}

func TestDynamicSizeListAcceptError(t *testing.T) {
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
		assert.Error(t, ts[0].Accept(ts[1]))
	}
}

func TestDynamicSizeListIsList(t *testing.T) {
	NewDynamicSizeList(NewNumber(""), "").isList()
}
