package bfhopefully_test

import (
	"bytes"
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
	if m.Operations[0] != b.IncrementPointer {
		t.Error("Unexpected Operation")
	}

	m = b.NewMachine([]b.Token{b.DecrementPointer, b.InputByte})
	if len(m.Operations) != 2 {
		t.Error("Unexpected Length")
	}
	if m.Operations[0] != b.DecrementPointer {
		t.Error("Unexpected Operation")
	}
	if m.Operations[1] != b.InputByte {
		t.Error("Unexpected Operation")
	}
}

func TestRun(t *testing.T) {
	input := []byte("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++.---.+++++++..+++.")
	r := bytes.NewReader(input)
	scanner := b.NewScanner(r)
	tokens, err := scanner.ScanAll()
		if err != nil {
		t.Error("Unexpected error %s", err)
	}
	m := b.NewMachine(tokens)
	output := m.Run()
	outputString := string(output)
	if outputString != "HELLO" {
		t.Error("Expecting 'HELLO' received " + outputString)
	}
}

func STestRun(t *testing.T) {
	input := []byte("++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++.")
	r := bytes.NewReader(input)
	scanner := b.NewScanner(r)
	tokens, err := scanner.ScanAll()
		if err != nil {
		t.Error("Unexpected error %s", err)
	}
	m := b.NewMachine(tokens)
	output := m.Run()
	outputString := string(output)
	if outputString != "Hello World!" {
		t.Error("Expecting 'Hello World!' received " + outputString)
	}
}
