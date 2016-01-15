package bfhopefully_test

import (
	"testing"
	b "bfhopefully"
	"bytes"
)

func TestScan(t *testing.T) {
	input := []byte("<>+- .,[]")
	r := bytes.NewReader(input)
	expectedTokens := []b.Token{b.DecrementPointer, b.IncrementPointer, b.IncrementByte, b.DecrementByte, b.Other, b.OutputByte, b.InputByte, b.JumpForward, b.JumpBackward}
	scanner := b.NewScanner(r)
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