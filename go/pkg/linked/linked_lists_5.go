package linked

type stack struct {
	ns []*Node
}

// RightToLeftSum is O(m + n)
func RightToLeftSum(l1, l2 List) List {
	n1 := l1.Head
	n2 := l2.Head
	// assumes length of each list is at least 1
	nextValue := n1.Value + n2.Value
	sumNode := &Node{Value: nextValue % 10, Next: nil}
	sumList := List{Head: sumNode}
	carry := nextValue / 10

	for n1.Next != nil || n2.Next != nil {
		if n1.Next == nil {
			nextValue = n2.Next.Value + carry
			n2 = n2.Next
		} else if n2.Next == nil {
			nextValue = n1.Next.Value + carry
			n1 = n1.Next
		} else {
			nextValue = n1.Next.Value + n2.Next.Value + carry
			n1 = n1.Next
			n2 = n2.Next
		}

		carry = nextValue / 10
		sumNode.Next = &Node{Value: nextValue % 10, Next: nil}
		sumNode = sumNode.Next
	}

	if carry == 1 {
		sumNode.Next = &Node{Value: 1, Next: nil}
	}

	return sumList
}

// O(n)
func (l List) listToStack() stack {
	stack := stack{}
	node := l.Head
	for node != nil {
		stack.ns = append(stack.ns, node)
		node = node.Next
	}
	return stack
}

// for both stack methods it is important to pass s by reference (*stack)

// O(1)
func (s *stack) len() int {
	return len(s.ns)
}

// O(1)
func (s *stack) pop() *Node {
	node := s.ns[len(s.ns)-1]
	s.ns = s.ns[0 : len(s.ns)-1]
	return node
}

// LeftToRightSum is O(m + n)
func LeftToRightSum(l1, l2 List) List {
	stack1 := l1.listToStack()
	stack2 := l2.listToStack()

	var head *Node
	var node *Node
	var nextValue int
	carry := 0

	for stack1.len() > 0 || stack2.len() > 0 {
		if stack1.len() == 0 {
			node = stack2.pop()
			nextValue = node.Value + carry
			node.Value = nextValue % 10
			node.Next = head
		} else if stack2.len() == 0 {
			node = stack1.pop()
			nextValue = node.Value + carry
			node.Value = nextValue % 10
			node.Next = head
		} else {
			node1 := stack1.pop()
			node2 := stack2.pop()
			nextValue = node1.Value + node2.Value + carry
			node = &Node{Value: nextValue % 10, Next: head}
		}

		carry = nextValue / 10
		head = node
	}

	if carry == 1 {
		node := &Node{Value: 1, Next: head}
		head = node
	}

	return List{Head: head}
}
