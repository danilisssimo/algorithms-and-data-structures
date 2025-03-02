# 🧱 Стек FIFO (First In, First Out)

## 📝 Описание
Стек FIFO (First In, First Out) — это структура данных, в которой элементы добавляются в конец и удаляются из начала. Это означает, что первый добавленный элемент будет первым удаленным. Такая структура часто используется в задачах, где важно сохранять порядок элементов, например, в очередях задач или буферах.

## 🛠 Основные операции

- **📥 Enqueue (Добавление в очередь)**: Добавляет элемент в конец стека.
- **📤 Dequeue (Удаление из очереди)**: Удаляет и возвращает элемент из начала стека.
- **🔍 GetLength (Длина очереди)**: Возвращает число элементов в стеке.
- **Draw (Вывод)**: Выводит очередь в консоль
- **GetNewQueue (Инициализация очереди)**: На вход принимает число (размер стека) и возвращает стек

## ⚠️ Примечания
- Время выполнения операции `dequeue` в наивной реализации (например, на основе списка) составляет O(n), так как удаление элемента из начала требует сдвига всех остальных элементов. Для более эффективной реализации рекомендуется использовать двусвязный список