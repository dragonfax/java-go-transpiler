package listen

import (
	"fmt"

	"github.com/dragonfax/java_converter/input/parser"
)

type StackListener struct {
	parser.JavaParserListener

	stack []parser.JavaParserListener
}

func NewStackListener() *StackListener {
	s := &StackListener{}
	s.stack = make([]parser.JavaParserListener, 0)
	return s
}

func (s *StackListener) Len() int {
	return len(s.stack)
}

func (s *StackListener) Peek() parser.JavaParserListener {
	return s.stack[len(s.stack)-1]
}

func (s *StackListener) Push(s2 parser.JavaParserListener) {
	s.stack = append(s.stack, s2)
	s.JavaParserListener = s2
}

func (s *StackListener) Pop(what parser.JavaParserListener) {
	if len(s.stack) == 1 {
		panic("can't pop the last element")
	}

	if s.JavaParserListener != what {
		panic(fmt.Sprintf("trying pop %T when %T", what, s.JavaParserListener))
	}

	if s.JavaParserListener != s.stack[len(s.stack)-1] {
		panic("current parser not last on list")
	}

	s.stack = s.stack[0 : len(s.stack)-1]
	s.JavaParserListener = s.stack[len(s.stack)-1]
}

type StackableListener struct {
	StackListener *StackListener
}

func NewStackableListener(sl *StackListener) *StackableListener {
	s := &StackableListener{StackListener: sl}
	return s
}

func (s *StackableListener) Pop(who interface{}, what parser.JavaParserListener) {
	s.StackListener.Pop(what)
}

func (s *StackableListener) Push(who interface{}, s2 parser.JavaParserListener) {
	s.StackListener.Push(s2)
}
