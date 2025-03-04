package linkedlist

func CreateLinkedList(slice []interface{}) *List {
	if len(slice) == 0 {
		return &List{
			HeadLink: nil,
			TailLink: nil,
			length:   len(slice),
		}
	}
	head := &ListNode{Data: slice[0]}
	currentNode := head
	for i := 1; i < len(slice); i++ {
		node := &ListNode{Data: slice[i]}
		currentNode.NextLink = node
		currentNode = currentNode.NextLink
	}
	return &List{
		HeadLink: head,
		TailLink: currentNode,
		length:   len(slice),
	}
}

// [1, 2, 3, 4, 5, 6, 7, 8, 9]
var List1 *List = CreateLinkedList([]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9})

// [15245, 16549879, 651987987, 6549844984, 65498767, 14897987, 875138, 498428]
var List2 *List = CreateLinkedList([]interface{}{15245, 16549879, 651987987, 6549844984, 65498767, 14897987, 875138, 498428})

// [42]
var List3 *List = CreateLinkedList([]interface{}{42})

// empty
var List4 *List = CreateLinkedList([]interface{}{})
