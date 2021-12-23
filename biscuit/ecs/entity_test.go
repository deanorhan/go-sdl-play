package ecs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEntity_New(t *testing.T) {
	e := NewEntity(nil)

	assert.Equal(t, uint32(0), e.id)
}

func TestEntity_AddComponent(t *testing.T) {
	e := new(Entity)
	c := new(Component)

	e.AddComponent(c)

	e.AddComponent(c)
	assert.Equal(t, 1, len(e.components))
}
