package bfhopefully

import (
	"testing"
	"bytes"
	"reflect"
)

func TestScan(t *testing.T) {
	input := []byte("<>+- .,[]")
	r := bytes.NewReader(input)
	expectedTokens := []Token{DecrementPointer, IncrementPointer, IncrementByte, DecrementByte, Other, OutputByte, InputByte, JumpForward, JumpBackward}
	scanner := NewScanner(r)
	for i := 0; i < len(input); i++ {
		token, err := scanner.Scan()
		if err != nil {
			t.Error("Unexpected error")
		}
		if token != expectedTokens[i] {
			t.Error("Unexpected token")
		}
	}
}

func TestScanAll(t *testing.T) {
	input := []byte("<>+- .,[]")
	r := bytes.NewReader(input)
	expectedTokens := []Token{DecrementPointer, IncrementPointer, IncrementByte, DecrementByte, Other, OutputByte, InputByte, JumpForward, JumpBackward}
	scanner := NewScanner(r)
	tokens, err := scanner.ScanAll()
	if err != nil {
		t.Error("Unexpected error %s", err)
	}
	if !reflect.DeepEqual(expectedTokens, tokens) {
		t.Error("ScanAll token mismatch")
	}
}
