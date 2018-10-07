package tinfer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunctionAccept(t *testing.T) {
	for _, ts := range [][2]Type{
		{
			NewFunction(nil, NewNumber(""), ""),
			NewFunction(nil, NewNumber(""), ""),
		},
		{
			NewFunction([]Type{NewNumber("")}, NewNumber(""), ""),
			NewFunction([]Type{NewNumber("")}, NewNumber(""), ""),
		},
		{
			NewFunction([]Type{NewNumber("")}, NewVariable(""), ""),
			NewFunction([]Type{NewNumber("")}, NewNumber(""), ""),
		},
		{
			NewFunction([]Type{NewVariable("")}, NewVariable(""), ""),
			NewFunction([]Type{NewVariable("")}, NewNumber(""), ""),
		},
	} {
		assert.Nil(t, ts[0].Accept(ts[1]))
	}
}

func TestFunctionAcceptError(t *testing.T) {
	for _, ts := range [][2]Type{
		{
			NewFunction(nil, NewNumber(""), ""),
			NewNumber(""),
		},
		{
			NewFunction(nil, NewNumber(""), ""),
			NewFunction(nil, NewFunction(nil, NewNumber(""), ""), ""),
		},
		{
			NewFunction([]Type{NewNumber("")}, NewNumber(""), ""),
			NewFunction([]Type{}, NewNumber(""), ""),
		},
		{
			NewFunction([]Type{NewNumber("")}, NewNumber(""), ""),
			NewFunction([]Type{NewFunction(nil, NewNumber(""), "")}, NewNumber(""), ""),
		},
	} {
		assert.Error(t, ts[0].Accept(ts[1]))
	}
}
