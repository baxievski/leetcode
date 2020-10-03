package removefromlinkedlist

import (
	"reflect"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	for _, test := range tests {
		expected := newLinkedList(test.expected)
		actual := RemoveNthFromEnd(newLinkedList(test.input), test.n)
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("\nRemoveNthFromEnd(%v, %v) -> %v, expected %v.", test.input, test.n, actual, expected)
		}
	}
}

func BenchmarkRemoveNthFromEnd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			RemoveNthFromEnd(newLinkedList(test.input), test.n)
		}
	}
}

func newLinkedList(values []int) *ListNode {
	dummy := ListNode{}
	c := &dummy
	for _, v := range values {
		c.Next = &ListNode{Val: v}
		c = c.Next
	}
	return dummy.Next
}
