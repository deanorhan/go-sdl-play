package ecs

import "reflect"

type System struct {
	components []*Component
}

func (s *System) AddComponent(comp *Component) {
	for _, component := range s.components {
		if reflect.TypeOf(comp) == reflect.TypeOf(component) {
			return
		}
	}

	s.components = append(s.components, comp)
}
