package nero

// Listener listens for actions of a specific type
type Listener interface {
	Listen(Action)
}

// ActionListener is used to register listeners of a given action type
type ActionListener struct {
	Action
	Listener
}
