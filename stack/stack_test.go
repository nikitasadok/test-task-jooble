package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack_NewStack(t *testing.T) {
	s := NewStack()
	assert.IsType(t, &Stack{}, s)
}

func TestStack_Push(t *testing.T) {
	s := NewStack()
	inputs := []string{"1", "2", "3"}
	outputs := []string{"3", "2", "1"}

	for _, inp := range inputs {
		s.Push(inp)
	}

	for _, out := range outputs {
		val := s.top.val
		assert.Equal(t, out, val)
		s.top = s.top.next
	}
}

func TestStack_Pop(t *testing.T) {
	testCases := []struct {
		stack *Stack
		want  []string
	}{
		{&Stack{}, []string{""}},
		{&Stack{top: &node{val: "1"}, size: 1}, []string{"1"}},
		{&Stack{top: &node{val: "1", next: &node{val: "2"}}, size: 2}, []string{"1", "2"}},
	}

	for _, testCase := range testCases {
		s := testCase.stack

		for _, want := range testCase.want {
			got := s.Pop()
			assert.Equal(t, want, got)
		}
	}
}

func TestStack_Peek(t *testing.T) {
	testCases := []struct {
		stack *Stack
		want  []string
	}{
		{&Stack{}, []string{""}},
		{&Stack{top: &node{val: "1"}, size: 1}, []string{"1"}},
		{&Stack{top: &node{val: "1", next: &node{val: "2"}}, size: 2}, []string{"1", "1"}},
	}

	for _, testCase := range testCases {
		s := testCase.stack

		for _, want := range testCase.want {
			got := s.Peek()
			assert.Equal(t, want, got)
		}
	}
}

func TestStack_IsEmpty(t *testing.T) {
	testCases := []struct {
		stack *Stack
		want  bool
	}{
		{&Stack{}, true},
		{&Stack{top: &node{val: "1"}, size: 1}, false},
	}
	for _, testCase := range testCases {
		s := testCase.stack
		isEmpty := s.IsEmpty()
		assert.Equal(t, isEmpty, testCase.want)
	}
}
