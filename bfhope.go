package bfhopefully

type Token int

const (
	Other Token = iota
	IncrementPointer
	DecrementPointer
	IncrementByte
	DecrementByte
	OutputByte
	InputByte
	JumpForward
	JumpBack
	
	IncrementPointerRune rune = '>'
	DecrementPointerRune rune = '<'
	IncrementByteRune rune = '+'
	DecrementByteRune rune = '-'
	OutputByteRune rune = '.'
	InputByteRune rune = ','
	JumpForwardRune rune = '['
	JumpBackwardRune rune = ']'
)

func IsIncrementPointer(ch rune) bool {
	return ch == IncrementPointerRune
}

func IsDecrementPointer(ch rune) bool {
	return ch == DecrementPointerRune
}

func IsIncrementByte(ch rune) bool {
	return ch == IncrementByteRune
}

func IsDecrementByte(ch rune) bool {
	return ch == DecrementByteRune
}

func IsOutputByte(ch rune) bool {
	return ch == OutputByteRune
}

func IsInputByte(ch rune) bool {
	return ch == InputByteRune
}

func IsJumpForward(ch rune) bool {
	return ch == JumpForwardRune
}

func IsJumpBackward(ch rune) bool {
	return ch == JumpBackwardRune
}

func IsOther(ch rune) bool {
	return !IsDecrementPointer(ch) && !IsDecrementPointer(ch) && !IsIncrementByte(ch) &&
		!IsDecrementByte(ch) && !IsOutputByte(ch) && !IsInputByte(ch) &&
		!IsJumpForward(ch) && !IsJumpBackward(ch)
}
