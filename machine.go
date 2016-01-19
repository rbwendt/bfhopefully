package bfhopefully

func incrementPointerFn() {}
func decrementPointerFn() {}
func incrementByteFn() {}
func decrementByteFn() {}
func outputByteFn() {}
func inputByteFn() {}
func jumpForwardFn() {}
func jumpBackwardFn() {}
	
var IncrementPointerFn = incrementPointerFn
var DecrementPointerFn = decrementPointerFn
var IncrementByteFn = incrementByteFn
var DecrementByteFn = decrementByteFn
var OutputByteFn = outputByteFn
var InputByteFn = inputByteFn
var JumpForwardFn = jumpForwardFn
var JumpBackwardFn = jumpBackwardFn

var TokenToOperation = map[Token]func()  {
	IncrementPointer: IncrementPointerFn,
}

type Machine struct {
	Position int
	State []byte
	Operations []func()
}

func NewMachine(t []Token) Machine {
	m := Machine{}
	for i:= 0; i < len(t); i++ {
		if t[i] == Other {
			continue
		}
		m.Operations = append(m.Operations, TokenToOperation[t[i]])
	}
	return m
}

