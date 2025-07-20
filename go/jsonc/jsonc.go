package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

const (
	LEFT_BRACE = iota
	RIGHT_BRACE
	COLON
	STRING
	BOOLEAN
	COMMA
	NUMBER
	NULL
	LEFT_BRACKET
	RIGHT_BRACKET
	EOF
)

type Token struct {
	Type    int
	Lexeme  string
	Literal interface{}
	Line    int
}

type Scanner struct {
	source  string
	tokens  []Token
	start   int
	current int
	line    int
	errors  []error
}

func NewScanner(source string) *Scanner {
	return &Scanner{
		source:  source,
		tokens:  []Token{},
		start:   0,
		current: 0,
		line:    1,
		errors:  []error{},
	}
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) peek() byte {
	return s.peekAt(0)
}

func (s *Scanner) peekNext() byte {
	return s.peekAt(1)
}

func (s *Scanner) advance() byte {
	c := s.source[s.current]
	s.current++
	if c == '\n' {
		s.line++
	}
	return c
}

func (s *Scanner) addTokenLiteral(tokenType int, literal interface{}) {
	text := s.source[s.start:s.current]
	s.tokens = append(s.tokens, Token{
		Type:    tokenType,
		Lexeme:  text,
		Literal: literal,
		Line:    s.line,
	})
}

func (s *Scanner) addToken(tokenType int) {
	s.addTokenLiteral(tokenType, nil)
}

func (s *Scanner) scanToken() error {
	s.start = s.current
	if s.isAtEnd() {
		s.addToken(EOF)
		return nil
	}
	c := s.advance()
	switch c {
	case ' ', '\r', '\t', '\n':
		// Ignore whitespace
	case '{':
		s.addToken(LEFT_BRACE)
	case '}':
		s.addToken(RIGHT_BRACE)
	case '[':
		s.addToken(LEFT_BRACKET)
	case ']':
		s.addToken(RIGHT_BRACKET)
	case ':':
		s.addToken(COLON)
	case ',':
		s.addToken(COMMA)
	case '"':
		s.string()
	default:
		if s.isDigit(c) || c == '-' {
			s.number()
		} else if s.isAlpha(c) {
			s.keyword()
		} else {
			s.error(fmt.Sprintf("Unexpected character: %c", c))
		}
	}
	return nil
}

func (s *Scanner) isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func (s *Scanner) keyword() error {
	for s.isAlpha(s.peek()) {
		s.advance()
	}

	text := s.source[s.start:s.current]
	switch text {
	case "true":
		s.addTokenLiteral(BOOLEAN, true)
	case "false":
		s.addTokenLiteral(BOOLEAN, false)
	case "null":
		s.addToken(NULL)
	default:
		s.error(fmt.Sprintf("Unexpected keyword: %s", text))
	}
	return nil
}

func (s *Scanner) string() error {
	// Skip opening quote, scan util closing quote
	for s.peek() != '"' && !s.isAtEnd() {
		if s.peek() == '\\' {
			s.advance()
			if !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.advance()
		}
	}
	if s.isAtEnd() {
		return s.error("Unterminated string")
	}
	s.advance()
	value := s.source[s.start+1 : s.current-1]
	s.addTokenLiteral(STRING, value)
	return nil
}

func (s *Scanner) number() error {

	if s.source[s.start] == '-' {
		if !s.isDigit(s.peek()) {
			return s.error("Invalid number")
		}
	}

	for s.isDigit((s.peek())) {
		s.advance()
	}

	if s.peek() == '.' && s.isDigit(s.peekNext()) {
		s.advance()
	}
	for s.isDigit(s.peek()) {
		s.advance()
	}
	value := s.source[s.start:s.current]
	num, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return s.error("Invalid number")
	}
	s.addTokenLiteral(NUMBER, num)
	return nil
}

func (s *Scanner) peekAt(n int) byte {
	pos := s.current + n
	if pos >= len(s.source) {
		return 0
	}
	return s.source[pos]
}

func (s *Scanner) isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func (s *Scanner) error(message string) error {
	err := fmt.Errorf("Error at line %d: %s\n", s.line, message)
	s.errors = append(s.errors, err)
	return err
}

func (s *Scanner) ScanTokens() (bool, error) {
	for !s.isAtEnd() {
		if err := s.scanToken(); err != nil {
			continue
		}
	}
	if len(s.errors) > 0 {
		return false, s.errors[0]
	}
	return true, nil
}

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}
	scanner := NewScanner(string(input))
	success, err := scanner.ScanTokens()
	if !success {
		fmt.Fprintf(os.Stderr, "Scanning failed: %v\n", err)
		// Print all errors
		for _, e := range scanner.errors {
			fmt.Fprintf(os.Stderr, "Error: %v\n", e)
		}
		os.Exit(1)
	}
	// success
	os.Exit(0)
}
