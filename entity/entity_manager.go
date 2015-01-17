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
	this.toAdd <- entity
}

func (this *EntityManager) Remove(entity *Entity) {
	this.toRemove <- entity
}

func (this *EntityManager) Update() {
	max := len(this.toAdd)
	for i := 0; i < max; i++ {
		this.entities = append(this.entities, <-this.toAdd)
	}

	max = len(this.toRemove)
	for i := 0; i < max; i++ {
		entity := <-this.toRemove
		for x := len(this.entities) - 1; x >= 0; x-- {
			if this.entities[x] == entity {
				this.entities = append(this.entities[:x], this.entities[x+1:]...)
			}
		}
	}

}
