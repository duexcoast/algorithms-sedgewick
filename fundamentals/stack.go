package fundamentals

import (
	"fmt"
	"strings"
)

type Stack struct {
	top *Node
	n   int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) IsEmpty() bool {
	return s.top == nil
}

func (s *Stack) Size() int {
	return s.n
}

func (s *Stack) Push(item Item) {
	s.top = newNode(item, s.top)
	s.n++
}

func (s *Stack) Pop() Item {
	if s.IsEmpty() {
		panic("stack is empty")
	}
	// remove item from top of stack
	item := s.top.item
	s.top = s.top.next
	s.n--
	return item
}

func (s *Stack) Peek() Item {
	if s.IsEmpty() {
		panic("stack is empty")
	}
	return s.top.item

}

func (s *Stack) Iterator() Iterator {
	items := make([]interface{}, s.n)

	i := 0
	cur := s.top
	for i < s.n {
		items[i] = cur.item
		cur = cur.next
		i++
	}
	return Iterator(items)
}

func (s *Stack) String() string {
	var ss []string

	for _, v := range s.Iterator() {
		ss = append(ss, fmt.Sprint(v))
	}

	return "[" + strings.Join(ss, ", ") + "]"
}
