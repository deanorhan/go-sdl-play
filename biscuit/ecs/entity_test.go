package ecs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEntityManager_NewEntity(t *testing.T) {

	t.Run("add a new entity", func(t *testing.T) {
		em := new(EntityManager)
		e := em.NewEntity()

		assert.Equal(t, uint32(1), e.Id())
		assert.Equal(t, 1, len(em.entities))
	})

	t.Run("adding more than allowed entities fails", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("code did not panic")
			}
		}()

		em := new(EntityManager)
		em.id_counter = maxEntityIds
		em.NewEntity()
	})

	t.Run("adding should resurrect from grveyard", func(t *testing.T) {
		em := new(EntityManager)
		e1 := em.NewEntity()

		em.RemoveEntity(e1)
		e2 := em.NewEntity()

		assert.Equal(t, e1.id, e2.id)
		assert.Equal(t, 1, len(em.entities))
		assert.Equal(t, 0, len(em.graveyard))
	})
}

func TestEntityManager_RemoveEntity(t *testing.T) {
	t.Run("remove an entity", func(t *testing.T) {
		em := new(EntityManager)
		e := em.NewEntity()

		em.RemoveEntity(e)
		assert.Equal(t, 0, len(em.entities))
		assert.Equal(t, 1, len(em.graveyard))
	})
}
