package ecs

import (
	"sync/atomic"
)

var idIncrement uint64

type BaseEntity struct {
	id uint64
	parent *BaseEntity
	children []*BaseEntity
}

type IdentifierSlice []Identifier

type Identifier interface {
	ID() uint64
}

type BaseFace interface {
	GetBaseEntity() *BaseEntity
}

// Just returns a pointer to the entity itself
// This means all entities containing a BaseEntity now have a GetBaseEntity method,
// allowing single-interfaced system.Add functions
func (e *BaseEntity) GetBaseEntity() *BaseEntity {
	return e
}

func NewBaseEntity() BaseEntity {
	return BaseEntity{id: atomic.AddUint64(&idIncrement, 1)}
}

func NewBaseEntities(amount int) []BaseEntity {
	entities := make([]BaseEntity, amount)

	lastID := atomic.AddUint64(&idIncrement, uint64(amount))
	for i := 0; i < amount; i++ {
		entities[i].id = lastID - uint64(amount) + uint64(i) + 1
	}

	return entities
}

func (e BaseEntity) ID() uint64 {
	return e.id
}

func (e *BaseEntity) AppendChild(child *BaseEntity) {
	child.parent = e
	e.children = append(e.children, child)
}

func (e *BaseEntity) RemoveChild(child *BaseEntity) {
	delete := -1
	for i, v := range e.children {
		if v.ID() == child.ID() {
			delete = i
			break
		}
	}

	if delete >= 0 {
		e.children = append(e.children[:delete], e.children[delete+1:]...)
	}
}

func (e *BaseEntity) Children() []BaseEntity {
	ret := []BaseEntity{}
	for _, child := range e.children {
		ret = append(ret, *child)
	}
	return ret
}

func (e *BaseEntity) Descendents() []BaseEntity {
	return descendents([]BaseEntity{}, e, e)
}

func descendents(in []BaseEntity, this, top *BaseEntity) []BaseEntity {
	for _, child := range this.children {
		in = descendents(in, child, top)
	}

	if this.ID() == top.ID() {
		return in
	}

	return append(in, *this)
}

func (e *BaseEntity) Parent() *BaseEntity {
	return e.parent
}

func (is IdentifierSlice) Len() int {
	return len(is)
}

func (is IdentifierSlice) Less(i, j int) bool {
	return is[i].ID() < is[j].ID()
}

func (is IdentifierSlice) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}