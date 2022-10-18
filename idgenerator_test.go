package common

import (
	"reflect"
	"testing"
)

func TestNewIDGenerator(t *testing.T) {
	gen, err := NewIDGenerator()
	if err != nil {
		t.Errorf("NewIDGenerator was incorrect, got error: %s", err)
	}

	id, err := gen.NextID()
	if err != nil {
		t.Errorf("NewIDGenerator was incorrect, got error: %s", err)
	}

	if reflect.TypeOf(id).String() != "uint64" {
		t.Errorf("NewIDGenerator was incorrect, got error: %s", err)
	}
}
