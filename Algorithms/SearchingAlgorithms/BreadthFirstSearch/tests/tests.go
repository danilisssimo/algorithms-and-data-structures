package test

import (
	"main/tree"
)

type TestCase struct {
	TreeLink  *tree.Node
	Searching int
	Result    bool
}

// Тестовое дерево 1:
//     1
//    / \
//   2   3
//  / \
// 4   5

var Tree1 *tree.Node = &tree.Node{
	Value: 1,
	Left: &tree.Node{
		Value: 2,
		Left:  &tree.Node{Value: 4},
		Right: &tree.Node{Value: 5},
	},
	Right: &tree.Node{Value: 3},
}

// Тестовое дерево 2:
//	   10
//	  /  \
//	 5    15
//	/ \     \
// 3   7     20

var Tree2 *tree.Node = &tree.Node{
	Value: 10,
	Left: &tree.Node{
		Value: 5,
		Left:  &tree.Node{Value: 3},
		Right: &tree.Node{Value: 7},
	},
	Right: &tree.Node{
		Value: 15,
		Right: &tree.Node{Value: 20},
	},
}

// Тестовое дерево 3: Пустое дерево
var Tree3 *tree.Node = nil

// Тестовое дерево 4: Дерево с одним узлом
var Tree4 *tree.Node = &tree.Node{Value: 42}

var TestCases = []TestCase{
	{TreeLink: Tree1, Searching: 4, Result: true},    // 4 существует в дереве
	{TreeLink: Tree1, Searching: 3, Result: true},    // 3 существует в дереве
	{TreeLink: Tree1, Searching: 6, Result: false},   // 6 не существует в дереве
	{TreeLink: Tree2, Searching: 20, Result: true},   // 20 существует в дереве
	{TreeLink: Tree2, Searching: 10, Result: true},   // 10 (корень) существует в дереве
	{TreeLink: Tree2, Searching: 8, Result: false},   // 8 не существует в дереве
	{TreeLink: Tree3, Searching: 1, Result: false},   // Пустое дерево, ничего не найдено
	{TreeLink: Tree4, Searching: 42, Result: true},   // 42 существует в дереве
	{TreeLink: Tree4, Searching: 100, Result: false}, // 100 не существует в дереве
}
