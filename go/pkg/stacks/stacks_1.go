package stacks

import (
	"fmt"
	"strconv"
	"strings"
)

// MinStack is LIFO data structure
// supporting O(1) push, pop, and min
type MinStack struct {
	Top      *MinFrame
	MinFrame *MinFrame
}

// ToString provides a string representation for a stack
func (s *MinStack) ToString() string {
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
		min := "None"

		if s.MinFrame != nil {
			min = strconv.Itoa(s.MinFrame.Data)
		} else {
			min = "nil"
		}

		return fmt.Sprintf("%s; Min: %s", strings.Join(ss, " -> "), min)
	}

	return "empty stack"
}

// Push O(1)
func (s *MinStack) Push(data int) {
	frame := &MinFrame{Data: data, Next: s.Top}

	if s.MinFrame != nil {
		if data < s.MinFrame.Data {
			frame.PreviousMinFrame = s.MinFrame
			s.MinFrame = frame
		}
	} else {
		s.MinFrame = frame
	}

	s.Top = frame
}

// Pop O(1)
func (s *MinStack) Pop() int {
	if s.Top == nil {
		panic("stack is empty")
	}

	if s.Top == s.MinFrame {
		s.MinFrame = s.Top.PreviousMinFrame
	}

	frame := s.Top
	s.Top = s.Top.Next

	return frame.Data
}

// Min O(1)
func (s *MinStack) Min() int {
	if s.Top == nil {
		panic("stack is empty")
	}

	return s.MinFrame.Data
}

// MinFrame represents a single stack frame for a MinStack
// using int for the data type out of convenience
type MinFrame struct {
	Data             int
	Next             *MinFrame
	PreviousMinFrame *MinFrame
}

// ToString provides a string representation for a stack frame
func (f *MinFrame) ToString() string {
	return fmt.Sprintf("[ %d ]", f.Data)
}
