package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	// check head isn't nil (null)
	if head == nil {
		return nil
	}

	// two pointers head, and next node. algorithm is drop next node if it's equal to the first ptr's value.
	for ptr1, ptr2 := head, head.Next; ptr2 != nil; ptr2 = ptr2.Next {
		// equal
		if ptr1.Val == ptr2.Val {
			// drop next value
			ptr1.Next = ptr2.Next
			// set second pointer to be traversed to new next node
			ptr2 = ptr1
		} else {
			// traverse ptr1, since it's sorted if the next node is different, there are no more similar nodes.
			ptr1 = ptr1.Next
		}
	}
	// pointer to head
	return head

}
