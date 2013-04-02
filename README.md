# fsm

## Table of Contents

0. [License](#license)
1. [Introduction](#introduction)
2. [Usage](#usage)
3. [Documentation] (#documentation)

## License

Copyright © 2000 Arturo Vergara
This work is free. You can redistribute it and/or modify it under the
terms of the Do What The Fuck You Want To Public License, Version 2,
as published by Sam Hocevar. See the COPYING file for more details.

## Introduction

`fsm` provides a simple finite state machine intended for game development.

## Usage

To use `fsm`, you first need to declare your own constant state types. For example:

```go

// states.go

package main

import "github.com/ArturoVM/fsm"

const (
    StateOne fsm.StateType = iota
    StateTwo
    StateThree
    StateQuit
)

```

After that, it's only a matter of writing your own states, conforming to the `fsm.GameState` interface (see the [documentation](#documentation) for more information) and initializing a package-wide `\*fsm.StateMachine` instance, and you're good to go:

```go

package main

import "github.com/ArturoVM/fsm"

type SomeState struct {
   stateType fsm.StateType
}

func (s *SomeState) Init() {
   s.stateType = StateOne
}

func (s *SomeState) Update() {
    // do some complicated logic
}

func (s *SomeState) Draw() {
    // do some drawing
}

// Game loop

var FSM *fsm.StateMachine // package-wide StateMachine instance

func m_init() {
    var someState SomeState
    FSM.Init(&someState)
}

func update() {
    FSM.Update() // Crucial! See documentation.
    FSM.CurrentState().Update()
}

func draw() {
    FSM.CurrentState().Draw()
}

func main() {
    m_init()
    for FSM.CurrentState().Type() != StateQuit {
        update()
        draw()
    }
}

```

## Documentation

You can find detailed documentation for this package, [here](http://godoc.org/github.com/ArturoVM/fsm)
