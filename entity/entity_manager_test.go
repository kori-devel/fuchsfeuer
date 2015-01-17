package entity

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEntityManagerCreate(t *testing.T) {

	Convey("Given a new entity manager", t, func() {
		manager := NewEntityManager()
		So(manager, ShouldNotBeNil)
		Convey("Check attributes", func() {
			So(manager.entities, ShouldNotBeNil)
			So(manager.toAdd, ShouldNotBeNil)
			So(manager.toRemove, ShouldNotBeNil)
		})

	})
}

func TestEntityManagerAdd(t *testing.T) {
	Convey("Given an entity manager and an entity", t, func() {
		manager := NewEntityManager()
		entity := New()
		Convey("Add entity", func() {
			manager.Add(entity)

			So(entity, ShouldBeIn, manager.entities)
		})
	})
}
