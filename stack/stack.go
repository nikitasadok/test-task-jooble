package stack

type node struct {
	val  string
	next *node
}

type Stack struct {
	top  *node
	size int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) Push(v string) {
	node := node{val: v}
	node.next = s.top
	s.top = &node
	s.size++
}

func (s *Stack) Pop() string {
	if s.IsEmpty() {
		return ""
	}

	v := s.top.val
	s.top = s.top.next
	s.size--
	return v
}

func (s *Stack) Peek() string {
	if s.IsEmpty() {
		return ""
	}
	return s.top.val
}
