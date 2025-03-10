# Установка и сброс бита

<pre>
<b>Уровень сложности</b>: Средний  
<b>Тема</b>: Побитовые операции  
</pre>

---

## Описание задачи  
Напишите функцию, которая устанавливает или сбрасывает бит в заданном числе.

### Входные данные:  
- `number` — целое число (0 ≤ number ≤ 2³¹ - 1).  
- `position` — номер бита (0 ≤ position ≤ 31).  
- `value` — 1 для установки бита, 0 для сброса.

### Выходные данные:  
- `int` — новое число после операции.

---

## Примеры

<p><strong class="example">Пример 1:</strong></p>
<pre><strong>Вход:</strong> number = 5, position = 1, value = 1
<strong>Выход:</strong> 7
</pre>

<p><strong class="example">Пример 2:</strong></p>
<pre><strong>Вход:</strong> number = 7, position = 0, value = 0
<strong>Выход:</strong> 6
</pre>