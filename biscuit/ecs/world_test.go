package ecs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorld_AddSystem(t *testing.T) {
	s := new(System)
	w := new(World)

	w.AddSystem(s)

	w.AddSystem(s)
	assert.Equal(t, 1, len(w.systems))
}

func TestWorld_RemoveSystem(t *testing.T) {
	s := new(System)
	w := new(World)

	w.AddSystem(s)

	w.RemoveSystem(s)
}

func TestWorld_AddEntity(t *testing.T) {
	e := new(Entity)
	w := new(World)

	w.AddEntity(e)

	w.AddEntity(e)
	assert.Equal(t, 1, len(w.entities))

	e2 := new(Entity)
	e2.id = 2
	w.AddEntity(e2)
	assert.Equal(t, 2, len(w.entities))
}
