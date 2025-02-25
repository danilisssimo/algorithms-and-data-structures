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

// Test Tree №1
//       1
//    /     \
//   2       3
//  / \     / \
// 4   5   6   7

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

// Test Tree №2
//     10
//    /  \
//   5   15
//  /      \
// 2       20

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

// Test Tree №3
//
//	     1
//	    / \
//	   2   5
//	  /     \
//	 3       6
//	/         \
// 4           7

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

// Test Tree №4
//       8
//    /     \
//   3       10
//  / \     /  \
// 1   6   9   14

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

// Test Tree №5
//      1
//    /   \
//   2     3
//    \   /
//     4 5

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

// Test Tree №6
//     1
//    / \
//   2   3
//  / \
// 4   5

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
var Tree8 *Node = nil
