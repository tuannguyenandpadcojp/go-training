package utils

import (
	"reflect"
	"testing"
)

func AssertCorrectResult[T any](t testing.TB, result, expected T) {
	t.Helper()
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected '%v' but got result '%v'", expected, result)
	}
}
