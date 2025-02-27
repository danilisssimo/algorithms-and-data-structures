package linkedlist

type ListValue struct {
	Data     interface{}
	NextLink *ListNode
}

type ListNode struct {
	Value    ListValue
	HeadLink *ListNode
	TailLink *ListNode
}

func (node *ListNode) CreateValue(data interface{}) {}
func (node *ListNode) AddToTop(value ListValue)     {}
func (node *ListNode) AddToEnd(value ListValue)     {}
func (node *ListNode) Remove()                      {}
