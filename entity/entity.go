package entity

import "fmt"

type Part interface{}

type Entity struct {
	components map[string]Part
}

func New() (this *Entity) {
	this = new(Entity)

	this.components = make(map[string]Part)

	return this
}

func (this *Entity) Attach(component Part, name string) (err error) {

	_, exists := this.components[name]
	if exists {
		return fmt.Errorf("Entity already has part with name %s", name)
	}

	this.components[name] = component

	return nil
}

func (this *Entity) Detach(name string) {
	delete(this.components, name)
}

func (this *Entity) Receive(name string) (part Part, err error) { return nil, nil }
