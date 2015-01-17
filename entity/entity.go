package entity

type Part interface{}

type Entity struct {
	components map[string]Part
}

func New() (this *Entity) {
	this = new(Entity)

	this.components = make(map[string]Part)

	return this
}

func (this *Entity) Attach(component Part, name string) {}

func (this *Entity) Detach(name string) {}
