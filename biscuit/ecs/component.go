package ecs

type Component interface {
	onUpdate()
	onRemove()
}

type ComponentManger struct {
	components map[uint32]*Component
}

func (cm *ComponentManger) initComponentMap() {
	if cm.components == nil {
		cm.components = make(map[uint32]*Component)
	}
}

func (cm *ComponentManger) AddComponent(e *Entity, c *Component) {
	cm.initComponentMap()

	cm.components[e.id] = c
}
