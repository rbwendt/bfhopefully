package bfhopefully

import (
	"testing"
)

func TestIsIncrementPointer(t *testing.T) {
	if !IsIncrementPointer(IncrementPointerRune) {
		t.Error("Expected > to be an IncrementPointer")
	}
	if IsIncrementPointer('p') {
		t.Error("Did not expected p to be an IncrementPointer")
	}
}

func TestIsDecrementPointer(t *testing.T) {
	if !IsDecrementPointer(DecrementPointerRune) {
		t.Error("Expected > to be a DecrementPointer")
	}
	if IsDecrementPointer('p') {
		t.Error("Did not expected p to be a DecrementPointer")
	}
}

func TestIsIncrementByte(t *testing.T) {
	if !IsIncrementByte(IncrementByteRune) {
		t.Error("Expected + to be an IncrementByte")
	}
	if IsIncrementByte('p') {
		t.Error("Did not expected p to be an IncrementByte")
	}
}

func TestIsDecrementByte(t *testing.T) {
	if !IsDecrementByte(DecrementByteRune) {
		t.Error("Expected - to be a DecrementByte")
	}
	if IsIncrementByte('p') {
		t.Error("Did not expected p to be a IncrementByte")
	}
}

func TestIsOutputByte(t *testing.T) {
	if !IsOutputByte(OutputByteRune) {
		t.Error("Expected . to be an OutputByte")
	}
	if IsOutputByte('p') {
		t.Error("Did not expected p to be an OutputByte")
	}
}

func TestIsJumpForward(t *testing.T) {
	if !IsJumpForward(JumpForwardRune) {
		t.Error("Expected . to be a JumpForward")
	}
	if IsJumpForward('p') {
		t.Error("Did not expected p to be a JumpForward")
	}
}

func TestIsJumpBackward(t *testing.T) {
	if !IsJumpBackward(JumpBackwardRune) {
		t.Error("Expected . to be a JumpBackward")
	}
	if IsJumpBackward('p') {
		t.Error("Did not expected p to be a JumpBackward")
	}
}

func TestIsOther(t *testing.T) {
	if IsOther(JumpBackwardRune) {
		t.Error("Expected . to be a JumpBackward")
	}
	if !IsOther('p') {
		t.Error("Expected . to be an Other")
	}

}
