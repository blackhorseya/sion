package app

// Servicer is a application service interface
type Servicer interface {
	// Start serve caller to start service
	Start() error

	// AwaitSignal serve caller to await signal
	AwaitSignal() error
}
