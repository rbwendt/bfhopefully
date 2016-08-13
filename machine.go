package bfhopefully

import (
	"bytes"
)

func (m *Machine) IncrementPointer() {
	if m.Position == len(m.State)-1 {
		m.State = append(m.State, byte(0))
	}
	m.Position++
}

func (m *Machine) DecrementPointer() {
	if m.Position == 0 {
		m.State = append([]byte{byte(0)}, m.State...)
	} else {
		m.Position--
	}
}

func (m *Machine) IncrementByte() {
	m.State[m.Position]++
}

func (m *Machine) DecrementByte() {
	m.State[m.Position]--
}

func (m *Machine) OutputByte() {
	m.Output = append(m.Output, m.State[m.Position])
}

func (m *Machine) InputByte() {
	input, _ := m.Input.ReadByte()
	m.State[m.Position] = byte(input)
}

func (m *Machine) JumpForward() {
	if m.State[m.Position] == 0 {
		parenCount := 1
		for i := m.Operation + 1; i < len(m.Operations); i++ {
			if m.Operations[i] == JumpBackward {
				parenCount--
				if parenCount == 0 {
					m.Operation = i
					break
				}
			} else if m.Operations[i] == JumpForward {
				parenCount++
			}
		}
	}
}

func (m *Machine) JumpBackward() {
	if m.State[m.Position] != 0 {
		parenCount := 1
		jumped := false
		for i := m.Operation - 1; i > 0; i-- {
			if m.Operations[i] == JumpForward {
				parenCount--
				if parenCount == 0 {
					m.Operation = i
					jumped = true
					break
				}
			} else if m.Operations[i] == JumpBackward {
				parenCount++
			}
		}
		if !jumped {
			m.Operation++
		}
	}
}

func (m *Machine) Run() []byte {
	for m.Operation < len(m.Operations) {
		currentOperation := m.Operations[m.Operation]
		if currentOperation == IncrementPointer {
			m.IncrementPointer()
		} else if currentOperation == DecrementPointer {
			m.DecrementPointer()
		} else if currentOperation == IncrementByte {
			m.IncrementByte()
		} else if currentOperation == DecrementByte {
			m.DecrementByte()
		} else if currentOperation == InputByte {
			m.InputByte()
		} else if currentOperation == OutputByte {
			m.OutputByte()
		} else if currentOperation == JumpForward {
			m.JumpForward()
		} else if currentOperation == JumpBackward {
			m.JumpBackward()
		}
		m.Operation++
	}
	return m.Output
}

type Machine struct {
	Position   int
	Operation  int
	State      []byte
	Output     []byte
	Input      *bytes.Reader
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
