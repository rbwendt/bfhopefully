package bfhopefully_test

import (
	"testing"
	b "github.com/rbwendt/bfhopefully/bfhopefully"
)

func TestIsIncrementPointer(t *testing.T) {
	if !b.IsIncrementPointer(b.IncrementPointerRune) {
		t.Error("Expected > to be an IncrementPointer")
	}
	if b.IsIncrementPointer('p') {
		t.Error("Did not expected p to be an IncrementPointer")
	}
}

func TestIsDecrementPointer(t *testing.T) {
	if !b.IsDecrementPointer(b.DecrementPointerRune) {
		t.Error("Expected > to be a DecrementPointer")
	}
	if b.IsDecrementPointer('p') {
		t.Error("Did not expected p to be a DecrementPointer")
	}
}

func TestIsIncrementByte(t *testing.T) {
	if !b.IsIncrementByte(b.IncrementByteRune) {
		t.Error("Expected + to be an IncrementByte")
	}
	if b.IsIncrementByte('p') {
		t.Error("Did not expected p to be an IncrementByte")
	}
}

func TestIsDecrementByte(t *testing.T) {
	if !b.IsDecrementByte(b.DecrementByteRune) {
		t.Error("Expected - to be a DecrementByte")
	}
	if b.IsIncrementByte('p') {
		t.Error("Did not expected p to be a IncrementByte")
	}
}

func TestIsOutputByte(t *testing.T) {
	if !b.IsOutputByte(b.OutputByteRune) {
		t.Error("Expected . to be an OutputByte")
	}
	if b.IsOutputByte('p') {
		t.Error("Did not expected p to be an OutputByte")
	}
}

func TestIsJumpForward(t *testing.T) {
	if !b.IsJumpForward(b.JumpForwardRune) {
		t.Error("Expected . to be a JumpForward")
	}
	if b.IsJumpForward('p') {
		t.Error("Did not expected p to be a JumpForward")
	}
}

func TestIsJumpBackward(t *testing.T) {
	if !b.IsJumpBackward(b.JumpBackwardRune) {
		t.Error("Expected . to be a JumpBackward")
	}
	if b.IsJumpBackward('p') {
		t.Error("Did not expected p to be a JumpBackward")
	}
}

func TestIsOther(t *testing.T) {
	if b.IsOther(b.JumpBackwardRune) {
		t.Error("Expected . to be a JumpBackward")
	}
	if !b.IsOther('p') {
		t.Error("Expected . to be an Other")
	}
	
}
