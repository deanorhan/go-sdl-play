package event

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvent_GetDispatcher(t *testing.T) {
	d := Get()

	assert.NotNil(t, d)
	assert.Equal(t, global_dispatch, d)
}

func TestDispatcher_Subscribe(t *testing.T) {
	d := Get()
	evt := Event{Name: "test"}

	d.Subscribe(evt, func(evt Event) {})
	assert.Equal(t, 1, len(global_dispatch.subscribers))

	err := global_dispatch.subscribe(evt, func() {})
	assert.Error(t, err)
}

func TestDispatcher_Fire(t *testing.T) {
	d := Get()
	evt := Event{Name: "test"}

	d.Subscribe(evt, func(evt Event) {})
	d.Fire(evt)
}
