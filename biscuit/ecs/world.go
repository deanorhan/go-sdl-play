package ecs

import (
	"reflect"
)

type World struct {
	systems  []*System
	entities []*Entity
}

func (w *World) AddSystem(sys *System) {
	for _, system := range w.systems {
		if reflect.TypeOf(sys) == reflect.TypeOf(system) {
			return
		}
	}

	w.systems = append(w.systems, sys)
}

func (w *World) RemoveSystem(sys *System) {
	delete := -1
	for idx, system := range w.systems {
		if reflect.TypeOf(sys) == reflect.TypeOf(system) {
			delete = idx
			break
		}
	}

	if delete >= 0 {
		w.systems = append(w.systems[:delete], w.systems[delete+1:]...)
	}
}

func (w *World) AddEntity(ent *Entity) {
	for _, entity := range w.entities {
		if entity.id == ent.id {
			return
		}
	}

	w.entities = append(w.entities, ent)
}
