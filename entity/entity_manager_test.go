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
