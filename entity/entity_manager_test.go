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

func TestEntityManagerAddRemove(t *testing.T) {
	Convey("Given an entity manager and an entity", t, func() {
		manager := NewEntityManager()
		entity := New()
		Convey("Add entity", func() {
			manager.Add(entity)
			manager.Update()

			So(entity, ShouldBeIn, manager.entities)
			Convey("And remove it", func() {
				manager.Remove(entity)
				manager.Update()

				So(entity, ShouldNotBeIn, manager.entities)
			})
		})
	})
}
