package entity

type EntityManager struct {
	toAdd    chan *Entity
	toRemove chan *Entity
	entities []*Entity
}

func NewEntityManager() (this *EntityManager) {
	return nil
}
