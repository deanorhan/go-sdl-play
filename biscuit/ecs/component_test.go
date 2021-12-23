package ecs

import "testing"

func TestComonentManager_AddComponent(t *testing.T) {
	cm := new(ComponentManger)
	e := new(Entity)
	c := new(Component)

	cm.AddComponent(e, c)
}
