# Алгоритмы и структуры данных

## Описание
Этот проект — моя личная инициатива по реализации алгоритмов, описанных в книге Томаса Кормена *«Алгоритмы. Построение и анализ»*. Основная цель проекта — углубить моё понимание того, как работают различные алгоритмы, их структуры и производительности. Проект не предназначен для практической пользы широкой аудитории, он создан исключительно для моего собственного удовольствия и обучения.

## Эффективность алгоритма
Одно из основных свойств алгоритма - его эффективность. Эффективность алгоритма показывает сколько вычислительных ресурсов потребуется для его выполнения. В наше время (на момент написания данной статьи) это может показаться не таким важным аспектом, так как современные компьютеры имеют огромные мощности, по сравнению с тем, что было 20 лет назад. Но вместе с данными мощностями растут и потребности: данных становится больше, задачи сложнее. Для решения данных задач потребуются ресурсы. И то, насколько эффективно мы можем использовать данные ресурсы, экономит компании огромную кучу денег.<br>
Эффективность алгоритма обеспечивается тем, сколько ресурсов он потребляет. Понять это можно посмотрев на его времнную сложность и пространственную сложность.

**Временная сложность** - это количество времени, которое потребуется алгоритму для выполнения задачи. Чем меньше временная сложность, тем быстрее алгоритм.<br>
**Пространственная сложность** - это объем памяти, который алгоритм будет использовать во время выполнения задачи. Эффективные алгоритмы стараются минимизировать объем потребляемой памяти. 

На самом деле эффективность алгоритма образуют куда больше параметров, но для моих задач этого будет достаточно.

Временная и пространственная сложности выражаются в виде функции от размера входных данных.

|  Обозначение  |        Название         |
| :-----------: | :---------------------: |
|    $O(1)$     |       постоянное        |
|  $O(log(n))$  |     логарифмическое     |
|    $O(n)$     |        линейное         |
| $O(nlog(n))$  | логарифмически линейное |
|   $O(n^2)$    |       квадратное        |
| $O(c^n), c>1$ |    экспоненциальное     |


## Задачи

Теория — это фундамент, но без практики она остаётся просто знаниями. Алгоритмы становятся по-настоящему понятными, когда мы применяем их на практике, решая различные задачи. Практика помогает:

- Улучшить понимание работы алгоритмов на реальных примерах.
- Научиться анализировать эффективность различных подходов.
- Развить навык оптимизации кода и выбора лучшего алгоритма для конкретной задачи.
- Подготовиться к техническим собеседованиям, где часто проверяют знание алгоритмов и структур данных.

В этом разделе будут собраны задачи, которые я решаю в процессе изучения алгоритмов. Они помогут закрепить теоретические знания и наработать интуицию в использовании различных методов.


## Реализованные алгоритмы и структуры данных

Этот список содержит алгоритмы, которые я изучил и реализовал. Он будет регулярно пополняться по мере моего обучения. Каждый алгоритм сопровождается описанием и разбором в соответствующем разделе.  

### Сортировка (Sorting Algorithms)  
- ✅ [Сортировка вставками (Insertion Sort)](/Algorithms/SortingAlgorithms/insertionSort/README.md)  
- ✅ [Сортировка слиянием (Merge Sort)](/Algorithms/SortingAlgorithms/mergeSort/README.md)  
- ✅ [Сортировка выбором (Selection Sort)](/Algorithms/SortingAlgorithms/selectionSort/README.md)  
- ✅ [Быстрая сортировка (Quick Sort)](/Algorithms/SortingAlgorithms/quickSort/README.md)  
- ⏳ Timsort (в процессе)  
- ⏳ Сортировка кучей (Heap Sort)  

### Поиск (Searching Algorithms)  
- ✅ [Линейный поиск (Linear Search)](/Algorithms/SearchingAlgorithms/LinearSearch/README.md)  
- ✅ [Бинарный поиск (Binary Search)](/Algorithms/SearchingAlgorithms/BinarySearch/README.md)  
- ✅ [Поиск в глубину (Depth-First Search)](/Algorithms/SearchingAlgorithms/DepthFirstSearch/README.md)  
- ✅ [Поиск в ширину (Breadth-First Search)](/Algorithms/SearchingAlgorithms/BreadthFirstSearch/README.md)

### 🏗 Структуры данных (Data Structures) 

#### 🌲 Деревья и графы
- ✅ [Двоичное дерево поиска (Binary Search Tree, BST)](/DataStructures/TreeStructure/BST/README.md)

#### Стопки (Stack)
- ✅ [FIFO (First In First Out)](/DataStructures/Stack/FIFO/README.md)

#### Списки (Lists)
- ✅ [Связный список (Linked List)](/DataStructures/Lists/LinkedList/README.md)


📢 **Планы**: В дальнейшем я добавлю больше алгоритмов, включая жадные алгоритмы, динамическое программирование, теорию графов и другие важные концепции.