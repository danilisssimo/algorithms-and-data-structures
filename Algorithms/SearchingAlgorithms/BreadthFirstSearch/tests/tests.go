package tests

import (
	tree "algorithms-and-data-structures/Algorithms/SearchingAlgorithms/DepthFirstSearch/tree"
	tests "algorithms-and-data-structures/DataStructures/TreeStructure/Tree"
)

type TestCase struct {
	TreeLink  *tree.Node
	Searching int
	Result    bool
}

var TestCases = []TestCase{
	{TreeLink: &tree.Node{Node: tests.Tree1}, Searching: 4, Result: true},    // 4 существует в дереве
	{TreeLink: &tree.Node{Node: tests.Tree1}, Searching: 3, Result: true},    // 3 существует в дереве
	{TreeLink: &tree.Node{Node: tests.Tree1}, Searching: 6, Result: false},   // 6 не существует в дереве
	{TreeLink: &tree.Node{Node: tests.Tree2}, Searching: 20, Result: true},   // 20 существует в дереве
	{TreeLink: &tree.Node{Node: tests.Tree2}, Searching: 10, Result: true},   // 10 (корень) существует в дереве
	{TreeLink: &tree.Node{Node: tests.Tree2}, Searching: 8, Result: false},   // 8 не существует в дереве
	{TreeLink: &tree.Node{Node: tests.Tree3}, Searching: 1, Result: false},   // Пустое дерево, ничего не найдено
	{TreeLink: &tree.Node{Node: tests.Tree4}, Searching: 42, Result: true},   // 42 существует в дереве
	{TreeLink: &tree.Node{Node: tests.Tree4}, Searching: 100, Result: false}, // 100 не существует в дереве
}
