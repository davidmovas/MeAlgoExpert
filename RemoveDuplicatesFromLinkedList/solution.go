package RemoveDuplicatesFromLinkedList

type ListNode struct {
	value int
	next  *ListNode
}

func RemoveDuplicatesFromLinkedList(node *ListNode) *ListNode {
	var current *ListNode
	head := node
	current = node

	for current.next != nil {
		if current.value == current.next.value {
			current.next = current.next.next
			continue
		}
		current = current.next
	}

	return head
}
