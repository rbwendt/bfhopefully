package bfhopefully

import (
	"bytes"
	"testing"
)

func TestNewMachine(t *testing.T) {

	m := NewMachine([]Token{Other})
	if len(m.Operations) != 0 {
		t.Error("Unexpected Operations")
	}

	m = NewMachine([]Token{IncrementPointer})
	if len(m.Operations) != 1 {
		t.Error("Unexpected Length")
	}
	if m.Operations[0] != IncrementPointer {
		t.Error("Unexpected Operation")
	}

	m = NewMachine([]Token{DecrementPointer, InputByte})
	if len(m.Operations) != 2 {
		t.Error("Unexpected Length")
	}
	if m.Operations[0] != DecrementPointer {
		t.Error("Unexpected Operation")
	}
	if m.Operations[1] != InputByte {
		t.Error("Unexpected Operation")
	}
}

func TestRunSimple(t *testing.T) {
	tokenInput := []byte("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++.---.+++++++..+++.")
	r := bytes.NewReader(tokenInput)
	scanner := NewScanner(r)
	tokens, err := scanner.ScanAll()
	if err != nil {
		t.Error("Unexpected error %s", err)
	}
	m := NewMachine(tokens)
	output := m.Run()
	outputString := string(output)
	if outputString != "HELLO" {
		t.Error("Expecting 'HELLO' received " + outputString)
	}
}

func TestRunWithInput(t *testing.T) {
	tokenInput := []byte(",.,.,.")
	r := bytes.NewReader(tokenInput)
	scanner := NewScanner(r)
	tokens, err := scanner.ScanAll()
	if err != nil {
		t.Error("Unexpected error %s", err)
	}
	m := NewMachine(tokens)
	m.Input = bytes.NewReader([]byte("YES"))
	output := m.Run()
	outputString := string(output)
	if outputString != "YES" {
		t.Error("Expecting 'YES' received " + outputString)
	}
}

func TestCommentLoop(t *testing.T) {
	tokenInput := []byte("[a comment loop.],.,.,.")
	r := bytes.NewReader(tokenInput)
	scanner := NewScanner(r)
	tokens, err := scanner.ScanAll()
	if err != nil {
		t.Error("Unexpected error %s", err)
	}
	m := NewMachine(tokens)
	m.Input = bytes.NewReader([]byte("YES"))
	output := m.Run()
	outputString := string(output)
	if outputString != "YES" {
		t.Error("Expecting 'YES' received " + outputString)
	}
}

func TestReadTooMuch(t *testing.T) {
	tokenInput := []byte(",.,.")
	r := bytes.NewReader(tokenInput)
	scanner := NewScanner(r)
	tokens, err := scanner.ScanAll()
	if err != nil {
		t.Error("Unexpected error %s", err)
	}
	m := NewMachine(tokens)
	m.Input = bytes.NewReader([]byte("Y"))
	output := m.Run()
	expected := append([]byte("Y"), byte(0))
	expectedString := string(expected)
	if string(output) != expectedString {
		t.Error("Expecting 'Y' received " + string(output))
	}

}

func TestRunLoop1(t *testing.T) {
	program := "--->->->>+>+>>+[++++[>+++[>++++>-->+++<<<-]<-]<+++]>>>.>-->-.>..+>++++>+++.+>-->[>-.<<]"
	input := []byte(program)
	r := bytes.NewReader(input)
	scanner := NewScanner(r)
	tokens, err := scanner.ScanAll()
	if err != nil {
		t.Error("Unexpected error %s", err)
	}
	m := NewMachine(tokens)
	output := m.Run()
	outputString := string(output)
	if outputString != "Hello, World!" {
		t.Error("Expecting 'Hello, World!' received " + outputString)
	}
}

func TestInterpreterBugs(t *testing.T) {
	program := ">++++++++[-<+++++++++>]<.>>+>-[+]++>++>+++[>[->+++<<+++>]<<]>-----.>->+++..+++.>-.<<+[>[+>+]>>]<--------------.>>.+++.------.--------.>+.>+."
	input := []byte(program)
	r := bytes.NewReader(input)
	scanner := NewScanner(r)
	tokens, err := scanner.ScanAll()
	if err != nil {
		t.Error("Unexpected error %s", err)
	}
	m := NewMachine(tokens)
	output := m.Run()
	outputString := string(output)
	expected := "Hello World!\n"
	if outputString != expected {
		t.Error("Expecting 'Hello, World!' received " + outputString)
	}
}

func TestRunHelloWorld1(t *testing.T) {
	program := "++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++."

	input := []byte(program)
	r := bytes.NewReader(input)
	scanner := NewScanner(r)
	tokens, err := scanner.ScanAll()
	if err != nil {
		t.Error("Unexpected error %s", err)
	}
	m := NewMachine(tokens)
	output := m.Run()
	outputString := string(output)
	expected := "Hello World!\n"
	if outputString != expected {
		t.Error("Expecting 'Hello World!' received " + outputString)
	}
}

func TestRunHelloWorld2(t *testing.T) {
	program := "++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>."

	input := []byte(program)
	r := bytes.NewReader(input)
	scanner := NewScanner(r)
	tokens, err := scanner.ScanAll()
	if err != nil {
		t.Error("Unexpected error %s", err)
	}
	m := NewMachine(tokens)
	output := m.Run()
	outputString := string(output)
	expected := "Hello World!\n"
	if outputString != expected {
		t.Error("Expecting 'Hello World!' received " + outputString)
	}
}
