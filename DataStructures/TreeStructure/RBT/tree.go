package rbt

type RBTNode struct {
	Value   interface{}
	IsBlack bool
	Left    *RBTNode
	Right   *RBTNode
}

func (node *RBTNode) Append(value interface{}) {}
func (node *RBTNode) Find(value interface{})   {}
func (node *RBTNode) Delete()                  {}

func (node *RBTNode) LeftRotation()  {}
func (node *RBTNode) RightRotation() {}
func (node *RBTNode) Recolor()       {}
