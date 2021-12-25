package ecs

import (
	"math"
	"sync/atomic"
)

const maxEntityIds = math.MaxUint32

type (
	Entity struct {
		id uint32
	}

	EntityManager struct {
		id_counter uint32
		entities   []Entity
		graveyard  []Entity
	}
)

func (e *Entity) Id() uint32 {
	return e.id
}

func (em *EntityManager) NewEntity() (e Entity) {
	if len(em.graveyard) > 0 {
		var e_temp Entity
		e_temp, em.graveyard = em.graveyard[0], em.graveyard[1:]
		e.id = e_temp.id

	} else {
		if em.id_counter == maxEntityIds {
			panic("exceeded entity assignment")
		}

		atomic.AddUint32(&em.id_counter, 1)
		e.id = em.id_counter
	}

	em.entities = append(em.entities, e)
	return
}

func (em *EntityManager) RemoveEntity(e Entity) {
	delete := -1

	for idx, ent := range em.entities {
		if ent.id == e.id {
			delete = idx
			break
		}
	}

	if delete >= 0 {
		em.entities = append(em.entities[:delete], em.entities[delete+1:]...)
		em.graveyard = append(em.graveyard, e)
	}
}
