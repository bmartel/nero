package nero

import "log"

// New creates a nero instance
func New() *Nero {
	n := &Nero{
		make(chan ActionListener),
		make(chan Action),
		make(chan bool),
		make(map[string][]Listener),
	}
	go n.run()
	return n
}

// Nero acts as the hub for distributing system actions
type Nero struct {
	register  chan ActionListener
	emitter   chan Action
	close     chan bool
	listeners map[string][]Listener
}

// Listen adds a listener for a given action
func (n *Nero) Listen(action Action, listener Listener) {
	n.register <- ActionListener{action, listener}
}

// Emit sends an action to all registered listeners
func (n *Nero) Emit(action Action) {
	n.emitter <- action
}

// Close nero channels
func (n *Nero) Close() {
	n.close <- true
}

func (n *Nero) run() {
	for {
		select {
		case listener := <-n.register:
			listenerGroup, exists := n.listeners[listener.Type()]
			if !exists {
				listenerGroup = make([]Listener, 0)
				log.Println("nero: no listener group found for type: " + listener.Type() + ", creating group now")
			}
			n.listeners[listener.Type()] = append(listenerGroup, listener.Listener)
		case action := <-n.emitter:
			listenerGroup, exists := n.listeners[action.Type()]
			if exists {
				log.Println("nero: emitting to all listeners action: " + action.Type())
				for _, listener := range listenerGroup {
					go listener.Listen(action)
				}
			}
		case <-n.close:
			n.listeners = make(map[string][]Listener)
			close(n.register)
			close(n.emitter)
			close(n.close)
			log.Println("nero: closing channel")
			return
		}
	}
}
