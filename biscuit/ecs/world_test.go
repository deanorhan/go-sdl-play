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
