package ecs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSystem_AddCompnent(t *testing.T) {
	s := new(System)
	c := new(Component)

	s.AddComponent(*c)

	s.AddComponent(*c)
	assert.Equal(t, 1, len(s.components))
}
