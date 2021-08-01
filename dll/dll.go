// package dll implements a Doubly linked List, with few needed methods.
package dll

// Node is the fundamental element of Dll.
type Node struct {
	Value interface{}
	Next  *Node
	Prev  *Node
}

// Dll encapsulates the state of the Doubly Linked list and has the needed methods.
type Dll struct {
	head *Node
	tail *Node
}

// New returns a new instance of Dll.
func New() *Dll {
	return &Dll{}
}

// GetHead returns the head of the list.
func (l *Dll) GetHead() *Node {
	return l.head
}

// Prepend adds the passed value to the front of the list.
func (l *Dll) Prepend(value interface{}) *Node {
	node := &Node{Value: value}

	if l.head == nil {
		l.tail = node
		l.head = node
	} else {
		node.Next = l.head
		l.head.Prev = node
		l.head = node
	}

	return node
}

// DeleteNode deletes the node referenced by ref.
func (l *Dll) DeleteNode(ref *Node) {
	if ref == nil {
		return
	}

	if ref.Prev != nil {
		ref.Prev.Next = ref.Next

		if ref.Next != nil {
			ref.Next.Prev = ref.Prev
		}
	} else {
		// ref is the head of list.
		if ref.Next != nil {

			ref.Next.Prev = nil
			l.head = ref.Next
		} else {
			l.head = nil
		}
	}
}

// DeleteAndInsertAtHead does what it's name suggests.
func (l *Dll) DeleteAndInsertAtHead(ref *Node) {
	if ref == nil {
		return
	}

	l.DeleteNode(ref)
	l.Prepend(ref.Value)
}
