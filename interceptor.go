package wrecker

// Interceptor contains functions that receive (and can modify) a
// wrecker.Request before it is sent to the server.  The Wrecker instance
// maintains an array of Interceptors that are applied to every
// wrecker.Request in the order that they were assigned.
type Interceptor struct {
	WreckerRequest func(*Request) error
}
