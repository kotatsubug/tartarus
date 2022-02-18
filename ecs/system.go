package ecs

type System interface {
	// Updates system, should be invoked once per frame.
	// dt is duration since last update.
	Update(dt float32)

	// Remove an entity from the system.
	Remove(e BaseEntity)
}

type SystemAddByInterfacer interface {
	System

	AddByInterface(o Identifier)
}

type Prioritizer interface {
	// Indicates the order in which systems should be executed per iteration.
	// Higher is sooner. Default is 0.
	Priority() int
}

type Initializer interface {
	// New initializes a system.
	// May be used to init some values beforehand like storing a reference to the World.
	New(*World)
}

// Sortable list of "System", indexed on System.Priority()
type systems []System

func (s systems) Len() int {
	return len(s)
}

func (s systems) Less(i, j int) bool {
	var prio1, prio2 int

	if prior1, ok := s[i].(Prioritizer); ok {
		prio1 = prior1.Priority()
	}
	if prior2, ok := s[j].(Prioritizer); ok {
		prio2 = prior2.Priority()
	}

	return prio1 > prio2
}

func (s systems) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}