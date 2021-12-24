package ecs

type Component interface {
}

type ComponentManger struct {
	components map[uint32]Component
}

func (cm *ComponentManger) initComponentMap() {
	if cm.components == nil {
		cm.components = make(map[uint32]Component)
	}
}

func (cm *ComponentManger) AddComponent(e Entity, c Component) {
	cm.initComponentMap()

	cm.components[e.id] = c
}

func (cm *ComponentManger) GetComponent(e Entity) (c Component) {
	if c, ok := cm.components[e.id]; ok {
		return c
	}

	return nil
}
