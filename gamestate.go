package fsm

// GameState holds the basic structure of the type
// the machine operates on.
type GameState interface {
	Init()
	Update()
	Draw()
	Type() StateType
}

/*
StateType is the type returned in GameState's Type method. You should
define your own constants in package main (or any other package of your
liking), like so:

    const (
	    FirstState fsm.StateType = iota
	    SecondState
	    ThirdState
    )
*/
type StateType uint8
