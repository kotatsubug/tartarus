package gfx

import (
	"fmt"
	"tartarus.xyz/ecs"
)

type TransformComponent struct {
	Transform Transform
}

type RenderableEntity struct {
	*ecs.BaseEntity
	*TransformComponent
}

type RenderableSystem struct {
	entities []RenderableEntity
}

func (m *RenderableSystem) Add(ent *ecs.BaseEntity, transform *TransformComponent) {
    m.entities = append(m.entities, RenderableEntity{ent, transform})
}

func (m *RenderableSystem) Remove(ent ecs.BaseEntity) {
	var delete int = -1
	for index, entity := range m.entities {
		  if entity.ID() == ent.ID() {
			  delete = index
			  break
		  }
	}
	if delete >= 0 {
		  m.entities = append(m.entities[:delete], m.entities[delete+1:]...)
	}
}

func (m *RenderableSystem) Update(dt float32) {
    for _, entity := range m.entities {
		entity.GetTransformComponent().Transform.Pos[0] += ((0.1 * dt * dt / 2))
        fmt.Println("Entity", entity.ID(), "reports dt", dt, "and transform", entity.GetTransformComponent().Transform.Pos)
    }
}

func (sc *TransformComponent) GetTransformComponent() *TransformComponent {
	return sc
}
