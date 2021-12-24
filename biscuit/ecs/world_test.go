package ecs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorld_NewEntity(t *testing.T) {
	w := new(World)
	w.NewEntity()
}

type SystemTest struct {
}

func (st *SystemTest) Init(*World) {}

func (st *SystemTest) Process(delta float64) {}

func TestWorld_AddSystem(t *testing.T) {
	w := new(World)
	s := SystemTest{}

	w.AddSystem(&s)
	w.AddSystem(&s)
	assert.Equal(t, 1, len(w.systems))
}

func TestWorld_RemoveSystem(t *testing.T) {
	w := new(World)
	s := SystemTest{}

	w.AddSystem(&s)
	w.RemoveSystem(&s)
	assert.Equal(t, 0, len(w.systems))
}
