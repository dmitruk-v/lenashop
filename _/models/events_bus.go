package models

import (
	"fmt"
	"log"
)

type EventHandler func(data interface{})

type EventsBus struct {
	subscribers map[string][]EventHandler
}

func NewEventBus() *EventsBus {
	return &EventsBus{subscribers: make(map[string][]EventHandler)}
}

func (bus *EventsBus) Notify(event string, data interface{}) {
	handlers, ok := bus.subscribers[event]
	if !ok {
		log.Printf("[EventBus]: Unknown event: ", event)
	}
	for _, handler := range handlers {
		handler(data)
	}
}

func (bus *EventsBus) Subscribe(event string, handler EventHandler) {
	subs := bus.subscribers[event]
	subs = append(subs, handler)
	bus.subscribers[event] = subs
}

// ----------------------------------------------------------------------
type A struct {
	ebus *EventsBus
}

func (a *A) Foo() {
	a.ebus.Subscribe("bla", func(data interface{}) {
		fmt.Println(data)
	})
}

type B struct {
	ebus *EventsBus
}

func (b *B) Bar() {
	b.ebus.Notify("bla", "--- Valera ---")
}

func Bla() {
	ebus := NewEventBus()
	a := A{ebus: ebus}
	b := B{ebus: ebus}

	fmt.Println(a, b)
	a.Foo()
}
