package entity

import "fmt"

type Part interface{}

type Entity struct {
	components map[string]Part
	actions    chan func()
}

func New() (this *Entity) {
	this = new(Entity)

	this.components = make(map[string]Part)
	this.actions = make(chan func(), 256)

	return this
}

func (this *Entity) Attach(component Part, name string) {

	this.actions <- func() {
		this.components[name] = component
	}

}

func (this *Entity) Detach(name string) {
	this.actions <- func() {
		delete(this.components, name)
	}
}

func (this *Entity) Receive(name string) (part Part, err error) {

	if _, exists := this.components[name]; !exists {
		return nil, fmt.Errorf("Can't return Part with name %s", name)
	}

	return this.components[name], nil
}

func (this *Entity) Update() {

	max := len(this.actions)

	for i := 0; i < max; i++ {
		f := <-this.actions
		f()
	}

}
