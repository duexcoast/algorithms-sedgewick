package fundamentals

type Stack[T any] struct {
	top *node[T]
	len int
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) IsEmpty() bool {
	return s.top == nil
}

func (s *Stack[T]) Size() int {
	return s.len
}

func (s *Stack[T]) Push(val T) {
	oldtop := s.top
	s.top = newNode(val)
	s.top.next = oldtop
	s.len++
}

func (s *Stack[T]) Pop() T {
	if s.IsEmpty() {
		panic("stack is empty")
	}
	// remove item from top of stack
	val := s.top.val
	s.top = s.top.next
	s.len--
	return val
}

func (s *Stack[T]) Peek() T {
	if s.IsEmpty() {
		panic("stack is empty")
	}
	return s.top.val

}

func (s *Stack[T]) Iterator() []T {
	items := make([]T, s.len)

	i := 0
	cur := s.top
	for i < s.len {
		items[i] = cur.val
		cur = cur.next
		i++
	}
	return items
}
