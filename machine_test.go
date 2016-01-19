package bfhopefully_test

import (
	"testing"
	b "bfhopefully"
)

func TestNewMachine(t *testing.T) {

	m := b.NewMachine([]b.Token{b.Other})
	if len(m.Operations) != 0 {
		t.Error("Unexpected Operations")
	}
	
	m = b.NewMachine([]b.Token{b.IncrementPointer})
	if len(m.Operations) != 1 {
		t.Error("Unexpected Length")
	}

	m = b.NewMachine([]b.Token{b.DecrementPointer, b.InputByte})
	if len(m.Operations) != 2 {
		t.Error("Unexpected Length")
	}

}