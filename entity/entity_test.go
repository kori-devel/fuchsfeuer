package entity

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewEntity(t *testing.T) {

	Convey("Given a new entity instance", t, func() {
		ent := New()

		Convey("Entity should not be nil", func() {
			So(ent, ShouldNotBeNil)
		})

		Convey("Entities attributes should be initialized", func() {
			So(ent.components, ShouldNotBeNil)
			So(ent.actions, ShouldNotBeNil)
		})
	})
}

func TestAttach(t *testing.T) {

	type E struct {
		A int
		B string
	}

	Convey("Given some components", t, func() {

		var tests = []struct {
			in   E
			out  E
			name string
		}{
			{
				E{A: 1, B: "hi"},
				E{A: 2, B: "Huhu"},
				"Test1",
			},
			{
				E{A: 2, B: "Huhu"},
				E{A: 2, B: "Huhu"},
				"Test1",
			},
			{
				E{A: 3, B: "Hihi"},
				E{A: 3, B: "Hihi"},
				"Test2",
			},
		}

		ent := New()

		Convey("Try to attach every component", func() {

			for _, v := range tests {
				ent.Attach(v.in, v.name)

			}

			ent.Update()

			for _, v := range tests {

				So(ent.components[v.name], ShouldResemble, v.out)
			}
		})

	})

}

func TestAttachThreaded(t *testing.T) {

	type E struct {
		A int
		B string
	}

	Convey("Given some components", t, func() {

		var tests = []struct {
			in   E
			out  E
			name string
		}{
			{
				E{A: 1, B: "hi"},
				E{A: 1, B: "hi"},
				"Test1",
			},
			{
				E{A: 2, B: "Huhu"},
				E{A: 2, B: "Huhu"},
				"Test1",
			},
			{
				E{A: 3, B: "Hihi"},
				E{A: 3, B: "Hihi"},
				"Test2",
			},
		}

		ent := New()

		for _, v := range tests {
			ent.Attach(v.in, v.name)

			ent.Update()
		}
		for _, v := range tests {
			go Convey(fmt.Sprintf("Try to attach component: %v", v.in), t, func() {

				So(ent.components[v.name], ShouldResemble, v.out)
			})
		}

	})

}

func TestDetach(t *testing.T) {

	type E struct {
		A int
		B string
	}

	Convey("Given some objects in entity.components", t, func() {

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

		var tests = []string{
			"Test1",
			"Test2",
			"Test3",
		}

		ent := New()

		for _, v := range mocks {
			ent.Attach(v.obj, v.name)
		}

		for _, v := range tests {
			ent.Detach(v)
		}

		ent.Update()

		Convey("Remove them from entity", func() {
			for _, v := range tests {

				So(ent.components[v], ShouldBeNil)

			}
		})
	})

}

func TestReceive(t *testing.T) {

	type E struct {
		A int
		B string
	}
	e := &E{A: 1, B: "hi"}
	f := &E{A: 2, B: "ho"}
	Convey("Given some objects in entity.components", t, func() {

		var tests = []struct {
			in          *E
			out         *E
			in_name     string
			out_name    string
			expectError bool
		}{
			{
				e,
				e,
				"Test1",
				"Test1",
				false,
			},
			{
				f,
				f,
				"Test2",
				"Test2",
				false,
			},
			{
				&E{A: 3, B: "he"},
				nil,
				"Test3",
				"Test4",
				true,
			},
		}

		ent := New()

		for _, v := range tests {
			ent.Attach(v.in, v.in_name)
			ent.Update()
		}

		Convey("Retrieve parts from entity", func() {
			for _, v := range tests {
				res, err := ent.Receive(v.out_name)

				if res != nil {
					So(res, ShouldEqual, v.out)
				} else {
					So(res, ShouldBeNil)
				}
				So(err != nil, ShouldEqual, v.expectError)
			}
		})

	})

}

func TestReceiveThreaded(t *testing.T) {

	type E struct {
		A int
		B string
	}
	e := &E{A: 1, B: "hi"}
	f := &E{A: 2, B: "ho"}
	Convey("Given some objects in entity.components", t, func() {

		var tests = []struct {
			in          *E
			out         *E
			in_name     string
			out_name    string
			expectError bool
		}{
			{
				e,
				e,
				"Test1",
				"Test1",
				false,
			},
			{
				f,
				f,
				"Test2",
				"Test2",
				false,
			},
			{
				&E{A: 3, B: "he"},
				nil,
				"Test3",
				"Test4",
				true,
			},
		}

		ent := New()

		for _, v := range tests {
			ent.Attach(v.in, v.in_name)
			ent.Update()
		}

		for _, v := range tests {
			go Convey(fmt.Sprintf("Retrieve part %v from entity", v), t, func() {
				res, err := ent.Receive(v.out_name)

				if res != nil {
					So(res, ShouldEqual, v.out)
				} else {
					So(res, ShouldBeNil)
				}
				So(err != nil, ShouldEqual, v.expectError)
			})
		}

	})

}
