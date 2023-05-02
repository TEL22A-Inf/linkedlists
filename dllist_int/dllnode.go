package dllist_int

// DLLNode represents an element in a doubly linked list.
type DLLNode struct {
	Prev  *DLLNode
	Value int
	Next  *DLLNode
}

// NewDLLNode creates a new empty doubly linked list element.
func NewDLLNode() *DLLNode {
	e := &DLLNode{}
	e.Prev = e
	e.Next = e
	return e
}

// IsEmpty returns true if the element is empty.
func (e *DLLNode) IsEmpty() bool {
	return e.Prev == e && e.Next == e
}

// Insert inserts a new element after the current element.
func (e *DLLNode) Insert(value int) *DLLNode {
	newElement := &DLLNode{Prev: e, Value: value, Next: e.Next}
	e.Next.Prev = newElement
	e.Next = newElement
	return newElement
}

// Remove removes the current element from the list.
// If the list is empty, nothing happens.
func (e *DLLNode) Remove() {
	if e.IsEmpty() {
		return
	}
	e.Prev.Next = e.Next
	e.Next.Prev = e.Prev
}
