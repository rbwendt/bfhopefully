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

func TestRunSimple(t *testing.T) {
	tokenInput := []byte("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++.---.+++++++..+++.")
	r := bytes.NewReader(tokenInput)
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

func TestRunWithInput(t *testing.T) {
	tokenInput := []byte(",.")
	r := bytes.NewReader(tokenInput)
	scanner := b.NewScanner(r)
	tokens, err := scanner.ScanAll()
		if err != nil {
		t.Error("Unexpected error %s", err)
	}
	m := b.NewMachine(tokens)
	m.Input = bytes.NewReader([]byte("Y"))
	output := m.Run()
	outputString := string(output)
	if outputString != "Y" {
		t.Error("Expecting 'Y' received " + outputString)
	}
}

func STestRunLoop(t *testing.T) {
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
