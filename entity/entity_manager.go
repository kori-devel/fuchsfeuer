package entity

type EntityManager struct {
	toAdd    chan *Entity
	toRemove chan *Entity
	entities []*Entity
}

func NewEntityManager() (this *EntityManager) {
	this = new(EntityManager)

	this.entities = make([]*Entity, 0)
	this.toAdd = make(chan *Entity, 64)
	this.toRemove = make(chan *Entity, 64)

	return this
}

func (this *EntityManager) Add(entity *Entity) {
}
