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

func TestAttach(t *testing.T) {

	type E struct {
		A int
		B string
	}

	var tests = []struct {
		in          E
		out         E
		name        string
		expectError bool
	}{
		{
			E{A: 1, B: "hi"},
			E{A: 1, B: "hi"},
			"Test1",
			false,
		},
		{
			E{A: 2, B: "Huhu"},
			E{A: 1, B: "hi"},
			"Test1",
			true,
		},
		{
			E{A: 3, B: "Hihi"},
			E{A: 3, B: "Hihi"},
			"Test2",
			false,
		},
	}

	ent := New()

	for _, v := range tests {
		err := ent.Attach(v.in, v.name)

		if (err != nil) != v.expectError {
			t.Errorf("entity.Attach(...) => %t, expected %t", err != nil, v.expectError)
		}

		if ent.components[v.name] != v.out {
			t.Errorf("entity.Attach(...) => %v, expected %v", ent.components[v.name], v.out)
		}
	}

}

func TestDetach(t *testing.T) {

}
