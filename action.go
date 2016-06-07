package nero

// Action is a system signal with a type descriptor
type Action interface {
	Type() string
}
