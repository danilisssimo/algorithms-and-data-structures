package tree

func CreateTree(height int) *Node {
	if height == 0 {
		return nil
	}
	return &Node{
		Value: height,
		Left:  CreateTree(height - 1),
		Right: CreateTree(height - 1),
	}
}

var Tree1 *Node = &Node{
	Value: 1,
	Left: &Node{
		Value: 2,
		Left: &Node{
			Value: 4,
		},
		Right: &Node{
			Value: 5,
		},
	},
	Right: &Node{
		Value: 3,
		Left: &Node{
			Value: 6,
		},
		Right: &Node{
			Value: 7,
		},
	},
}

var Tree2 *Node = &Node{
	Value: 10,
	Left: &Node{
		Value: 5,
		Left: &Node{
			Value: 2,
		},
	},
	Right: &Node{
		Value: 15,
		Right: &Node{
			Value: 20,
		},
	},
}

var Tree3 *Node = &Node{
	Value: 1,
	Left: &Node{
		Value: 2,
		Left: &Node{
			Value: 3,
			Left: &Node{
				Value: 4,
			},
		},
	},
	Right: &Node{
		Value: 5,
		Right: &Node{
			Value: 6,
			Right: &Node{
				Value: 7,
			},
		},
	},
}

var Tree4 *Node = &Node{
	Value: 8,
	Left: &Node{
		Value: 3,
		Left: &Node{
			Value: 1,
		},
		Right: &Node{
			Value: 6,
		},
	},
	Right: &Node{
		Value: 10,
		Left: &Node{
			Value: 9,
		},
		Right: &Node{
			Value: 14,
		},
	},
}

var Tree5 *Node = &Node{
	Value: 1,
	Left: &Node{
		Value: 2,
		Right: &Node{
			Value: 4,
		},
	},
	Right: &Node{
		Value: 3,
		Left: &Node{
			Value: 5,
		},
	},
}

var Tree6 *Node = &Node{
	Value: 1,
	Left: &Node{
		Value: 2,
		Left:  &Node{Value: 4},
		Right: &Node{Value: 5},
	},
	Right: &Node{Value: 3},
}
var Tree7 *Node = CreateTree(15)
