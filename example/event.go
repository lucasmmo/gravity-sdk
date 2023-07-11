package main

import (
	"fmt"

	"github.com/lucasmmo/gravity-sdk/event"
)

// Event

type Event struct {
	key  string
	data []byte
}

func (e *Event) GetKey() string {
	return e.key
}
func (e *Event) GetData() interface{} {
	return e.data
}

// Usecase

type someStruct struct {
	dispatcher event.Dispatcher
}

func NewSomeStruct(dispatcher event.Dispatcher) *someStruct {
	return &someStruct{dispatcher}
}
func (s *someStruct) Something(data []byte) error {
	// Do something here
	s.dispatcher.Dispatch(&Event{
		key:  "some_event",
		data: data,
	})
	return nil
}

// Listener
type listener struct {
	data []byte
}

func NewListener() *listener {
	return &listener{}
}
func (l *listener) Handler() error {
	fmt.Println(l.data.(map[string][]byte})["gremio"])
	return nil
}
func (l *listener) SetData(data []byte) {
	l.data = data
}

func main() {

	d := event.NewDispatcher()
	l := NewListener()
	d.AddListener("some_event", l)

	s := NewSomeStruct(d)
	s.Something(map[string]interface{}{
		"message": "Hello world",
		"gremio":  "foda",
	})

}
