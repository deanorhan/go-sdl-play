package event

import (
	"errors"
	"sync"
)

type (
	Event struct {
		Name string
	}

	Subscriber interface {
		Handle(Event)
	}

	dispatcher struct {
		sync.RWMutex

		subscribers map[string][]interface{}
	}

	Dispatcher interface {
		Subscribe(Event, ...interface{})
		Fire(Event)
	}

	handlerFn = func(Event)
)

var global_dispatch = &dispatcher{
	subscribers: make(map[string][]interface{}),
}

func Get() Dispatcher {
	return global_dispatch
}

func (d *dispatcher) Subscribe(evt Event, subs ...interface{}) {
	d.Lock()
	defer d.Unlock()

	for _, sub := range subs {
		d.subscribe(evt, sub)
	}
}

func (d *dispatcher) subscribe(evt Event, sub interface{}) error {
	if _, ok := sub.(handlerFn); ok {
		d.subscribers[evt.Name] = append(d.subscribers[evt.Name], sub)
		return nil
	}

	return errors.New("subscriber must be a function")
}

func (d *dispatcher) Fire(evt Event) {
	d.RLock()
	defer d.RUnlock()

	subs := d.subscribers[evt.Name]
	for _, sub := range subs {
		d.call(evt, sub)
	}
}

func (d *dispatcher) call(evt Event, sub interface{}) {
	if handler, ok := sub.(handlerFn); ok {
		handler(evt)
	}
}
