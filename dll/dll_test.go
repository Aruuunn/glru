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

func assertAfterPrepend(t *testing.T, values []interface{}) {
	l := dll.New()

	for _, value := range values {
		l.Prepend(value)
	}

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
