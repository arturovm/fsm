// Package fsm provides a simple finite state machine intended for game
// development.
package fsm

// StateMachine is the machine itself.
// It holds information about the current state and the state it
// should move to next.
type StateMachine struct {
	currentState GameState
	nextState    GameState
}

// Init initializes an instance of StateMachine. You must provide the intial
// state.
func (m *StateMachine) Init(gs GameState) {
	m.currentState = gs
	m.currentState.Init()
	m.nextState = nil
}

/*
Update checks if the machine is ready to move on to the next state. If
it is, it will simply set its current state to the new one.

See SetNextState for more information
*/
func (m *StateMachine) Update() {
	if m.nextState != nil {
		m.currentState = m.nextState
		m.currentState.Init()
		m.nextState = nil
	}
}

// SetNextState tells the machine it should move from its current state,
// to the state you specify
func (m *StateMachine) SetNextState(gs GameState) {
	m.nextState = gs
}

// CurrentState returns the machine's current state. Useful for calling
// the state's Update and Draw methods.
func (m *StateMachine) CurrentState() GameState {
	return m.currentState
}
