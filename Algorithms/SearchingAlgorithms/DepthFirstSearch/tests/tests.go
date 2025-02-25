package tests

import (
	DFSTree "algorithms-and-data-structures/Algorithms/SearchingAlgorithms/DepthFirstSearch/tree"
	tests "algorithms-and-data-structures/DataStructures/TreeStructure/Tree"
)

type TestCase struct {
	TreeLink  *DFSTree.Node
	Searching int
	Result    bool
}

var TestCases = []TestCase{
	// Tree2
	{TreeLink: &DFSTree.Node{Node: tests.Tree2}, Searching: 10, Result: true},   // Корень дерева
	{TreeLink: &DFSTree.Node{Node: tests.Tree2}, Searching: 5, Result: true},    // Левый узел
	{TreeLink: &DFSTree.Node{Node: tests.Tree2}, Searching: 2, Result: true},    // Левый лист
	{TreeLink: &DFSTree.Node{Node: tests.Tree2}, Searching: 15, Result: true},   // Правый узел
	{TreeLink: &DFSTree.Node{Node: tests.Tree2}, Searching: 20, Result: true},   // Правый лист
	{TreeLink: &DFSTree.Node{Node: tests.Tree2}, Searching: 100, Result: false}, // Несуществующее значение

	// Tree3
	{TreeLink: &DFSTree.Node{Node: tests.Tree3}, Searching: 1, Result: true},  // Корень дерева
	{TreeLink: &DFSTree.Node{Node: tests.Tree3}, Searching: 2, Result: true},  // Левый узел
	{TreeLink: &DFSTree.Node{Node: tests.Tree3}, Searching: 3, Result: true},  // Левый лист
	{TreeLink: &DFSTree.Node{Node: tests.Tree3}, Searching: 4, Result: true},  // Левый лист
	{TreeLink: &DFSTree.Node{Node: tests.Tree3}, Searching: 5, Result: true},  // Правый узел
	{TreeLink: &DFSTree.Node{Node: tests.Tree3}, Searching: 6, Result: true},  // Правый лист
	{TreeLink: &DFSTree.Node{Node: tests.Tree3}, Searching: 7, Result: true},  // Правый лист
	{TreeLink: &DFSTree.Node{Node: tests.Tree3}, Searching: 8, Result: false}, // Несуществующее значение

	// Tree4
	{TreeLink: &DFSTree.Node{Node: tests.Tree4}, Searching: 8, Result: true},  // Корень дерева
	{TreeLink: &DFSTree.Node{Node: tests.Tree4}, Searching: 3, Result: true},  // Левый узел
	{TreeLink: &DFSTree.Node{Node: tests.Tree4}, Searching: 1, Result: true},  // Левый лист
	{TreeLink: &DFSTree.Node{Node: tests.Tree4}, Searching: 6, Result: true},  // Левый лист
	{TreeLink: &DFSTree.Node{Node: tests.Tree4}, Searching: 10, Result: true}, // Правый узел
	{TreeLink: &DFSTree.Node{Node: tests.Tree4}, Searching: 9, Result: true},  // Правый лист
	{TreeLink: &DFSTree.Node{Node: tests.Tree4}, Searching: 14, Result: true}, // Правый лист
	{TreeLink: &DFSTree.Node{Node: tests.Tree4}, Searching: 0, Result: false}, // Несуществующее значение

	// Tree5
	{TreeLink: &DFSTree.Node{Node: tests.Tree5}, Searching: 1, Result: true},   // Корень дерева
	{TreeLink: &DFSTree.Node{Node: tests.Tree5}, Searching: 2, Result: true},   // Левый узел
	{TreeLink: &DFSTree.Node{Node: tests.Tree5}, Searching: 4, Result: true},   // Левый лист
	{TreeLink: &DFSTree.Node{Node: tests.Tree5}, Searching: 3, Result: true},   // Правый узел
	{TreeLink: &DFSTree.Node{Node: tests.Tree5}, Searching: 5, Result: true},   // Правый лист
	{TreeLink: &DFSTree.Node{Node: tests.Tree5}, Searching: 10, Result: false}, // Несуществующее значение

	// Tree6
	{TreeLink: &DFSTree.Node{Node: tests.Tree6}, Searching: 1, Result: true},  // Корень дерева
	{TreeLink: &DFSTree.Node{Node: tests.Tree6}, Searching: 2, Result: true},  // Левый узел
	{TreeLink: &DFSTree.Node{Node: tests.Tree6}, Searching: 4, Result: true},  // Левый лист
	{TreeLink: &DFSTree.Node{Node: tests.Tree6}, Searching: 5, Result: true},  // Левый лист
	{TreeLink: &DFSTree.Node{Node: tests.Tree6}, Searching: 3, Result: true},  // Правый узел
	{TreeLink: &DFSTree.Node{Node: tests.Tree6}, Searching: 6, Result: false}, // Несуществующее значение

	// Tree8 (пустое дерево)
	{TreeLink: &DFSTree.Node{Node: tests.Tree8}, Searching: 1, Result: false}, // Пустое дерево

}
