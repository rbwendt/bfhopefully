package bfhopefully

import (
	"bufio"
)

func (m *Machine) IncrementPointerFn() {
	if m.Position == len(m.State) - 1 {
		m.State = append(m.State, byte(0))
	}
	m.Position ++
}

func (m *Machine) DecrementPointerFn() {
	if m.Position == 0 {
		m.State = append([]byte{byte(0)}, m.State...)
	} else {
		m.Position --
	}
}

func (m *Machine) incrementByteFn() {
	m.State[m.Position]++
}

func (m *Machine) decrementByteFn() {
	m.State[m.Position]--
}

func (m *Machine) outputByteFn() {
	m.Output = append(m.Output, m.State[m.Position])
}

func (m *Machine) InputByteFn() {
	input, _ := m.Input.ReadRune()
	m.State[m.Position] = byte(input)
}

func (m *Machine) JumpForwardFn() {
	if m.State[m.Position] == 0 {
		parenCount := 1
		for i := m.Operation; i < len(m.Operations) ; i++ {
			if m.Operations[i] == JumpBackward {
				parenCount --
				if parenCount == 0 {
					m.Operation = i
					break;
				}
			} else if m.Operations[i] == JumpForward {
				parenCount ++
			}
		}
	}
}

func (m *Machine) JumpBackwardFn() {
	if m.State[m.Position] != 0 {
		parenCount := 1
		for i := m.Operation; i > 0 ; i-- {
			if m.Operations[i] == JumpForward {
				parenCount --
				if parenCount == 0 {
					m.Operation = i
					break;
				}
			} else if m.Operations[i] == JumpBackward {
				parenCount ++
			}
		}
	}
}

type Machine struct {
	Position int
	Operation int
	State []byte
	Output []byte
	Input bufio.Reader
	Operations []Token
}

func NewMachine(t []Token) Machine {
	m := Machine{}
	m.State = append(m.State, byte(0))
	for i := 0; i < len(t); i++ {
		if t[i] != Other {
			m.Operations = append(m.Operations, t[i])
		}
	}
	return m
}

