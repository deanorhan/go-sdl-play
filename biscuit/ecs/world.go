package ecs

import (
	"reflect"
)

type World struct {
	systems  []System
	entities EntityManager
}

func NewWorld() (e *World) {
	return &World{
		entities: EntityManager{
			id_counter: 0,
		},
	}
}

func (w *World) NewEntity() (e Entity) {
	return w.entities.NewEntity()
}

func (w *World) RemoveEntity(e Entity) {
	w.entities.RemoveEntity(e)
}

func (w *World) AddSystem(sys System) {
	if init, ok := sys.(SystemInit); ok {
		init.Init(w)
	}

	for _, system := range w.systems {
		if reflect.TypeOf(sys) == reflect.TypeOf(system) {
			return
		}
	}

	w.systems = append(w.systems, sys)
}

func (w *World) RemoveSystem(sys System) {
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

func (w *World) Process(delta float64) {
	for _, system := range w.systems {
		system.Process(delta)
	}
}
