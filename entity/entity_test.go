package entity

import "testing"

func TestNewEntity(t *testing.T) {

	ent := New()

	if ent == nil {
		t.Errorf("New() => nil, expected entity")
	}

	if ent.components == nil {
		t.Errorf("New().components => nil, expected map")
	}

}
