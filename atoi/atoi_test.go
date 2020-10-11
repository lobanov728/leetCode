package main

import (
	"reflect"
	"testing"
)

func TestSimple(t *testing.T) {
	var assert int = 42

	res := MyAtoi("42")
	if !reflect.DeepEqual(assert, res) {
		t.Error(res, " not equal ", assert)
	}
}

func TestMinusSimple(t *testing.T) {
	var assert int = -42
	res := MyAtoi("   -42")
	if !reflect.DeepEqual(assert, res) {
		t.Error(res, " not equal ", assert)
	}
}

func TestWithWords(t *testing.T) {
	var assert int = 4193
	res := MyAtoi("4193 with words")
	if !reflect.DeepEqual(assert, res) {
		t.Error(res, " not equal ", assert)
	}
}