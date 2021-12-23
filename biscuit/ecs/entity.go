package ecs

import (
	"reflect"

	"github.com/veandco/go-sdl2/sdl"
)

type Entity struct {
	id         uint32
	components []*Component
}

type Component interface {
	onUpdate()
	onRemove()
}

func NewEntity(renderer *sdl.Renderer) (e Entity) {
	e.id = 0

	return e
}

func (e *Entity) AddComponent(comp *Component) {
	for _, component := range e.components {
		if reflect.TypeOf(comp) == reflect.TypeOf(component) {
			return
		}
	}

	e.components = append(e.components, comp)
}
