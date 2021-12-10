package main

import "github.com/dragonfax/delver_converter/parser"

type StackListener struct {
	parser.JavaParserListener

	stack []parser.JavaParserListener
}

func NewStackListener() *StackListener {
	s := &StackListener{}
	s.stack = make([]parser.JavaParserListener, 0)
	return s
}

func (s *StackListener) Push(s2 parser.JavaParserListener) {
	s.stack = append(s.stack, s2)
	s.JavaParserListener = s2
}

func (s *StackListener) Pop() {
	if len(s.stack) == 0 {
		panic("poping end of stack")
	}
	s.JavaParserListener = s.stack[len(s.stack)-1]
}
