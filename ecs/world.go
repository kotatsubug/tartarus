package ecs

import (
	"reflect"
	"sort"
)

// Contains a bunch of entities and a bunch of systems separately.
type World struct {
	systems systems
	sysIn, sysEx map[reflect.Type][]reflect.Type
}

// Adds a system to the World sorted by priority.
func (w *World) AddSystem(system System) {
	if initializer, ok := system.(Initializer); ok {
		initializer.New(w)
	}
	w.systems = append(w.systems, system)
	sort.Sort(w.systems)
}

// Adds a system to the world, but also adds a filter that:
// - allows automatic adding of entities that match the provided "in" interface
// - excludes any that match the provided "ex" interface (even if they match "in")
// "in" and "ex" MUST be pointers to the interface.
func (w *World) AddSystemInterface(sys SystemAddByInterfacer, in interface{}, ex interface{}) {
	w.AddSystem(sys)

	if w.sysIn == nil {
		w.sysIn = make(map[reflect.Type][]reflect.Type)
	}

	if !reflect.TypeOf(in).AssignableTo(reflect.TypeOf([]interface{}{})) {
		in = []interface{}{in}
	}

	for _, v := range in.([]interface{}) {
		w.sysIn[reflect.TypeOf(sys)] = append(w.sysIn[reflect.TypeOf(sys)], reflect.TypeOf(v).Elem())
	}

	if ex == nil {
		return
	}

	if w.sysEx == nil {
		w.sysEx = make(map[reflect.Type][]reflect.Type)
	}

	if !reflect.TypeOf(ex).AssignableTo(reflect.TypeOf([]interface{}{})) {
		ex = []interface{}{ex}
	}
	for _, v := range ex.([]interface{}) {
		w.sysEx[reflect.TypeOf(sys)] = append(w.sysEx[reflect.TypeOf(sys)], reflect.TypeOf(v).Elem())
	}
}

// Adds the entity to all systems that have been added using AddSystemInterface
// NOTE: If the system was added using AddSystem the entity will NOT be added to it.
func (w *World) AddEntity(e Identifier) {
	if w.sysIn == nil {
		w.sysIn = make(map[reflect.Type][]reflect.Type)
	}
	if w.sysEx == nil {
		w.sysEx = make(map[reflect.Type][]reflect.Type)
	}

	search := func(i Identifier, types []reflect.Type) bool {
		for _, t := range types {
			if reflect.TypeOf(i).Implements(t) {
				return true
			}
		}
		return false
	}
	for _, system := range w.systems {
		sys, ok := system.(SystemAddByInterfacer)
		if !ok {
			continue
		}

		if ex, not := w.sysEx[reflect.TypeOf(sys)]; not {
			if search(e, ex) {
				continue
			}
		}
		if in, ok := w.sysIn[reflect.TypeOf(sys)]; ok {
			if search(e, in) {
				sys.AddByInterface(e)
				continue
			}
		}
	}
}

// Get list of systems managed by the world.
func (w *World) Systems() []System {
	return w.systems
}

// Updates all systems managed by the World
func (w *World) Update(dt float32) {
	for _, system := range w.Systems() {
		system.Update(dt)
	}
}

// Removes entity across ALL systems
func (w *World) RemoveEntity(e BaseEntity) {
	for _, sys := range w.systems {
		sys.Remove(e)
	}
}

// Sorts the systems in the world by priority.
func (w *World) SortSystems() {
	sort.Sort(w.systems)
}