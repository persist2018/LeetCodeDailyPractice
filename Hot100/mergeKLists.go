package main

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return merge(lists, 0, len(lists)-1)
}

func merge(lists []*ListNode, left, right int) *ListNode {
	if left == right {
		return lists[left]
	}
	mid := left - (left-right)/2
	node1 := merge(lists, left, mid)
	node2 := merge(lists, mid+1, right)
	return mergeTwoSortedLists(node1, node2)
}

// This method just iterate mergeTwoSortedLists
func mergeKListsByMergingTwoSortedLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	ans := lists[0]
	for i := 1; i < len(lists); i++ {
		ans = mergeTwoSortedLists(ans, lists[i])
	}
	return ans
}

func mergeTwoSortedLists(Node1, Node2 *ListNode) *ListNode {
	if Node1 == nil {
		return Node2
	}
	if Node2 == nil {
		return Node1
	}
	head := new(ListNode)
	p1 := Node1
	p2 := Node2
	p := head
	for p1 != nil && p2 != nil {
		if p1.Val > p2.Val {
			p.Next = p2
			p2 = p2.Next
		} else {
			p.Next = p1
			p1 = p1.Next
		}
		p = p.Next
	}
	if p1 == nil {
		p.Next = p2
	}
	if p2 == nil {
		p.Next = p1
	}
	return head.Next
}
