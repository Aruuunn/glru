package dll_test

import (
	"log"
	"reflect"
	"testing"

	"github.com/ArunMurugan78/glru/dll"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	l := dll.New()
	assert.NotNil(t, l)
	assert.Equal(t, "<*dll.Dll Value>", reflect.ValueOf(l).String())
}

func TestPrepend(t *testing.T) {
	assertAfterPrepend(t, []interface{}{})
	assertAfterPrepend(t, []interface{}{1})
	assertAfterPrepend(t, []interface{}{"One", 2, "three", 6, 1.2})
}

func TestDeleteNode(t *testing.T) {
	assertAfterDeletion(t, 0, []interface{}{"One"})
	assertAfterDeletion(t, 1, []interface{}{1, 2, 3, 4})
	assertAfterDeletion(t, 3, []interface{}{1, 2, 3, 4})

	assertAfterDeletion(t, 0, []interface{}{1, 2, 3, 4})
}

func assertAfterDeletion(t *testing.T, index int, values []interface{}) {
	l := getListAfterPrepending(values)

	idx := 0
	ptr := l.GetHead()
	for idx < index {
		ptr = ptr.Next
		idx++
	}
	l.DeleteNode(ptr)
	assert.Equal(t, true, traverseAndCheckValue(l.GetHead(), remove(values, index)))
}

func getListAfterPrepending(values []interface{}) *dll.Dll {
	l := dll.New()

	for _, value := range values {
		l.Prepend(value)
	}

	return l
}

func assertAfterPrepend(t *testing.T, values []interface{}) {
	l := getListAfterPrepending(values)
	assert.Equal(t, true, traverseAndCheckValue(l.GetHead(), values))
}

func traverseAndCheckValue(head *dll.Node, expectedValues []interface{}) bool {
	if (head != nil) != (len(expectedValues) != 0) {
		return false
	}

	if head == nil {
		return true
	}

	var prev *dll.Node = nil

	for i := len(expectedValues) - 1; i >= 0; i-- {
		if head.Prev != prev {
			return false
		}

		if expectedValues[i] != head.Value {
			log.Printf("Expected %v but got %v\n", expectedValues[i], head.Value)
			return false
		}

		prev = head
		head = head.Next
	}

	return true
}

func TestTestUtil(t *testing.T) {
	head := &dll.Node{Value: "One"}

	head.Next = &dll.Node{Value: "Two"}

	head.Next.Prev = head

	assert.Equal(t, true, traverseAndCheckValue(head, []interface{}{"Two", "One"}))
}

func remove(slice []interface{}, s int) []interface{} {
	return append(slice[:(len(slice)-s-1)], slice[(len(slice)-s-1)+1:]...)
}
