package bfhopefully

import (
	"bufio"
	"io"
)

var eof = rune(0)

type Scanner struct {
	r *bufio.Reader
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

func (s *Scanner) read() (rune, error) {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof, err
	}
	return ch, nil
}

func (s *Scanner) unread() {
	_ = s.r.UnreadRune()
}

func (s *Scanner) Scan() (token Token, err error) {
	ch, err := s.read()
	if err != nil {
		return token, err
	}
	if IsIncrementPointer(ch) {
		token = IncrementPointer
	} else if IsDecrementPointer(ch) {
		token = DecrementPointer
	} else if IsIncrementByte(ch) {
		token = IncrementByte
	} else if IsDecrementByte(ch) {
		token = DecrementByte
	} else if IsOutputByte(ch) {
		token = OutputByte
	} else if IsInputByte(ch) {
		token = InputByte
	} else if IsJumpForward(ch) {
		token = JumpForward
	} else if IsJumpBackward(ch) {
		token = JumpBackward
	} else {
		token = Other
	}
	return token, nil
}

func (s *Scanner) ScanAll() (tokens []Token, err error) {
	scanning := true
	for scanning {
		token, err := s.Scan()
		if err == io.EOF {
			break
		}
		if err != nil {
			return tokens, err
		}
		tokens = append(tokens, token)
	}
	return tokens, nil
}