package event

type Dispatcher interface {
	AddListener(eventName string, listener Listener)
	Dispatch(event Event)
}

type Listener interface {
	SetData(data interface{})
	Handler() error
}

type Event interface {
	GetKey() string
	GetData() interface{}
}

type dispatcher struct {
	Listeners map[string][]Listener
}

func NewDispatcher() *dispatcher {
	return &dispatcher{
		Listeners: make(map[string][]Listener),
	}
}

func (d *dispatcher) AddListener(eventName string, listener Listener) {
	d.Listeners[eventName] = append(d.Listeners[eventName], listener)
}

func (d *dispatcher) Dispatch(event Event) {
	if d.Listeners == nil {
		return
	}
	for _, listener := range d.Listeners[event.GetKey()] {
		listener.SetData(event.GetData())
		listener.Handler()
	}
}
