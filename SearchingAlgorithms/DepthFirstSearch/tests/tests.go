package test

import (
	"main/tree"
)

type TestCase struct {
	TreeLink  *tree.TreeNode
	Searching int
	Result    bool
}

// Тестовое дерево 1:
//     1
//    / \
//   2   3
//  / \
// 4   5

var Tree1 *tree.TreeNode = &tree.TreeNode{
	Value: 1,
	Left: &tree.TreeNode{
		Value: 2,
		Left:  &tree.TreeNode{Value: 4},
		Right: &tree.TreeNode{Value: 5},
	},
	Right: &tree.TreeNode{Value: 3},
}

// Тестовое дерево 2:
//
//	   10
//	  /  \
//	 5    15
//	/ \     \
//
// 3   7     20

var Tree2 *tree.TreeNode = &tree.TreeNode{
	Value: 10,
	Left: &tree.TreeNode{
		Value: 5,
		Left:  &tree.TreeNode{Value: 3},
		Right: &tree.TreeNode{Value: 7},
	},
	Right: &tree.TreeNode{
		Value: 15,
		Right: &tree.TreeNode{Value: 20},
	},
}

// Тестовое дерево 3: Пустое дерево
var Tree3 *tree.TreeNode = nil

// Тестовое дерево 4: Дерево с одним узлом
var Tree4 *tree.TreeNode = &tree.TreeNode{Value: 42}

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
