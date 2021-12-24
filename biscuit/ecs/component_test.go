package ecs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComonentManager_AddComponent(t *testing.T) {
	cm := new(ComponentManger)
	e := new(Entity)
	c := new(Component)

	cm.AddComponent(*e, *c)
}

type MatchingComponent struct {
}

func TestComonentManager_GetComponent(t *testing.T) {
	cm := new(ComponentManger)
	c := &MatchingComponent{}
	e := Entity{id: 0}

	cm.AddComponent(e, c)
	c2 := cm.GetComponent(e)
	assert.Same(t, c, c2)
}
