package linked

// see Remove Duplicates for List and Node type definitions
// see Sum Lists for stack type definition and methods

// IsPalindrome is O(n)
func IsPalindrome(l List) bool {
	listNode := l.Head
	stack := l.ListToStack()
	listLength := stack.Len()

	for stack.Len() >= listLength/2 {
		stackNode := stack.Pop()
		if stackNode.Value != listNode.Value {
			return false
		}
		listNode = listNode.Next
	}

	return true
}
