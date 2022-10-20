package event_test

import (
	"testing"

	"github.com/lucasmmo/gravity-sdk/event"
)

type testListener struct {
	data   interface{}
	called bool
}

func (l *testListener) Handler() error {
	l.called = true
	return nil
}

func (l *testListener) SetData(data interface{}) {
	l.data = data
}

type testEvent struct {
	key  string
	data interface{}
}

func (t *testEvent) GetKey() string {
	return t.key
}

func (t *testEvent) GetData() interface{} {
	return t.data
}

const (
	EVENT_NAME = "test"
	EVENT_DATA = "message"
)

func TestDispatcher(t *testing.T) {
	t.Run("should add listener", func(t *testing.T) {
		ed := event.NewDispatcher()
		listener := &testListener{}

		ed.AddListener(EVENT_NAME, listener)

		if len(ed.Listeners) != 1 {
			t.Errorf("Expect: %v; Received: %v", len(ed.Listeners), 1)
		}
	})
	t.Run("should dispatch", func(t *testing.T) {
		ed := event.NewDispatcher()
		listener := &testListener{}
		event := &testEvent{EVENT_NAME, EVENT_DATA}

		ed.AddListener(EVENT_NAME, listener)

		ed.Dispatch(event)

		if !listener.called {
			t.Errorf("Expect: %v; Received: %v", true, listener.called)
		}
	})
}
