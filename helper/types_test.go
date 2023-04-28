package helper

import (
	"errors"
	"reflect"
	"testing"
)

// Should return correct type and implements interface
func TestMapError(t *testing.T) {
	err := errors.New("aybjax")
	var errResult interface{} = NewMapError(err)

	expected := &mapError{
		"error": err.Error(),
	}

	if !reflect.DeepEqual(errResult, expected) {
		t.Error("Incorrect type is returned")
	}

	if _, ok := errResult.(error); !ok {
		t.Error("Incorrect type is returned")
	}
}
