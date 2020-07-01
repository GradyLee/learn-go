package algorithm

type listNode struct {
	data 	int
	next 	*listNode
}

const bucketNum = 10

func listInsert(head *listNode, val int) (*listNode) {
	var newnode *listNode
	var dummyNode listNode
	curr := head
	pre := &dummyNode
	dummyNode.next = head

	for curr != nil && curr.data < val {
		pre, curr = curr, curr.next
	}

	newnode = new(listNode)
	newnode.data = val
	newnode.next = curr
	pre.next = newnode

	return dummyNode.next
}

func listMerge(head1, head2 *listNode) (*listNode) {
	pos1 := head1
	pos2 := head2
	var dummyNode listNode
	dummy := &dummyNode

	for pos1 != nil && pos2 != nil {
		if (pos1.data <= pos2.data) {
			dummy.next = pos1
			pos1 = pos1.next
		} else {
			dummy.next = pos2
			pos2 = pos2.next
		}
		dummy = dummy.next
	}

	if pos1 != nil {
		dummy.next = pos1
	}
	if pos2 != nil {
		dummy.next = pos2
	}

	return dummyNode.next
}

func BucketSort(v []int) {
	var buckets [bucketNum]*listNode
	for _, d := range v {
		i := d % bucketNum
		buckets[i] = listInsert(buckets[i], d)
	}

	for i := 1; i < bucketNum; i++ {
		buckets[0] = listMerge(buckets[0], buckets[i])
	}

	var index int
	for p := buckets[0]; p != nil; p = p.next {
		v[index] = p.data
		index++
	}
}