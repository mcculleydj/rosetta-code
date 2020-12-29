package stacks

import (
	"fmt"
	"strings"
)

// StackSet is LIFO data structure
// with multiple stacks
// supporting O(1) push, pop, and popAt
type StackSet struct {
	Stacks        []*Stack
	Index         int
	StackCapacity int
}

// ToString provides a string representation for a stack set
func (ss *StackSet) ToString() string {
	stacks := ""

	for i, s := range ss.Stacks {
		stacks += fmt.Sprintf("%d: %s\n", i, s.ToString())
	}

	return stacks
}

// Pop is O(1)
func (ss *StackSet) Pop() int {
	if len(ss.Stacks) == 0 {
		panic("stack set is empty")
	}

	for ss.Stacks[ss.Index].Top == nil {
		ss.Stacks = ss.Stacks[:len(ss.Stacks)-1]
		ss.Index--

		if ss.Index == -1 {
			ss.Index = 0
			panic("stack set is empty")
		}
	}

	return ss.Stacks[ss.Index].Pop()
}

// Push is O(1)
func (ss *StackSet) Push(data int) {
	if len(ss.Stacks) == 0 {
		ss.Stacks = append(ss.Stacks, &Stack{})
		ss.Push(data)
	} else if ss.Stacks[ss.Index].Length == ss.StackCapacity {
		ss.Index++
		ss.Stacks = append(ss.Stacks, &Stack{})
		ss.Push(data)
	} else {
		ss.Stacks[ss.Index].Push(data)
	}
}

// PopAt is O(1)
// this does not move frames between stacks
func (ss *StackSet) PopAt(i int) int {
	if i >= len(ss.Stacks) {
		panic("index out of bounds")
	}

	if ss.Stacks[i].Top == nil {
		panic("stack is empty")
	}

	return ss.Stacks[i].Pop()
}

// Stack is a LIFO data structure
type Stack struct {
	Top    *Frame
	Length int
}

// ToString provides a string representation for a stack
func (s *Stack) ToString() string {
	frames := []string{}
	top := s.Top

	for top != nil {
		frames = append(frames, top.ToString())
		top = top.Next
	}

	if len(frames) > 0 {
		ss := []string{"top"}
		ss = append(ss, frames...)
		ss = append(ss, "bottom")

		return strings.Join(ss, " -> ")
	}

	return "empty stack"
}

// Push O(1)
func (s *Stack) Push(data int) {
	frame := &Frame{Data: data, Next: s.Top}
	s.Top = frame
	s.Length++
}

// Pop O(1)
func (s *Stack) Pop() int {
	if s.Top == nil {
		panic("stack is empty")
	}

	frame := s.Top
	s.Top = s.Top.Next
	s.Length--

	return frame.Data
}

// Frame is a single stack frame
type Frame struct {
	Data int
	Next *Frame
}

// ToString provides a string representation for a stack frame
func (f *Frame) ToString() string {
	return fmt.Sprintf("[ %d ]", f.Data)
}
