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

	type E struct {
		A int
		B string
	}

	var mocks = []struct {
		obj  E
		name string
	}{
		{
			E{1, "hi"},
			"Test1",
		},
		{
			E{2, "ho"},
			"Test2",
		},
	}

	var tests = []struct {
		name         string
		expectRemove bool
	}{
		{
			"Test1",
			true,
		},
		{
			"Test2",
			true,
		},
		{
			"Test2",
			false,
		},
		{
			"Test3",
			false,
		},
	}

	ent := New()

	for _, v := range mocks {
		ent.Attach(v.obj, v.name)
	}

	for _, v := range tests {
		_, existsBefore := ent.components[v.name]
		ent.Detach(v.name)
		_, existsAfter := ent.components[v.name]

		if (existsBefore && !existsAfter) != v.expectRemove {
			t.Errorf("entity.Detach(...) => %t, expected %t", existsBefore && !existsAfter, v.expectRemove)
		}

	}

}

func TestReceive(t *testing.T) {

	type E struct {
		A int
		B string
	}

	var tests = []struct {
		in       *E
		out      *E
		in_name  string
		out_name string
	}{
		{
			&E{A: 1, B: "hi"},
			&E{A: 1, B: "hi"},
			"Test1",
			"Test1",
		},
		{
			&E{A: 1, B: "ho"},
			&E{A: 1, B: "ho"},
			"Test2",
			"Test2",
		},
		{
			&E{A: 3, B: "he"},
			nil,
			"Test3",
			"Test4",
		},
	}

	ent := New()

	for _, v := range tests {
		ent.Attach(v.in, v.in_name)

		res, _ := ent.Receive(v.out_name)

		if res != v.out {
			t.Errorf("entity.Recive(...) => %v, expected %v", res, v.out)
		}
	}

}
