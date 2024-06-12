---
layout: layouts/post.njk
title: Стек наиболее часто встречающихся элементов
date: 2024-06-03
complexity: hard
original_url: https://leetcode.com/problems/maximum-frequency-stack/description/
preview_image: /images/previews/maximum_frequency_stack.webp
tags:
  - hard
  - heap
  - hash table
---
---

## Описание задачи

Необходимо реализовать структуру данных, которая позволит сохранять и доставать целочисленные элементы.
Для этого необходимо реализовать 3 функции:

1. `Constructor() FreqStack` конструктор, инициализирующий структуры данных.
2. `FreqStack.Push(val int)` метод структуры данных, позволяющий сохранить целочисленный элемент внутрь структуры
   данных `FreqStack`
3. `FreqStack.Pop() int` метод структуры данных, удаляющий и возвращающий элемент из структуры данных, по следующему
   принципу:

- Если в `FreqStack` существует элемент, который встречается чаще остальных, из стек извлекается этот элемент (последний
  добавленный из серии).
- Если в `FreqStack` существуют несколько наборов элементов с одинаковой частотой, извлекается элемент из того набора,
  чей элемент был добавлен а `FreqStack` последним.
- Если в `FreqStack` все элементы имеют одинаковую частоту/все элементы встречаются только 1 раз, элемент извлекается по
  принципу стека.

---

## Ограничения

- Добавляемые/извлекаемые элементы лежат в диапазоне от 0 до 1 000 000 000
- Максимальное количество вызовов функций `Push` и `Pop` - 20 000
- Гарантируется, что перед вызовом функции `Pop` в `FreqStack` будет как минимум один элемент

---

## Примеры

{% tabs %}
{% tab "Пример №1" %}

**Инициализация**:

```go
stack := Constructor()

stack.Push(5)
stack.Push(7)
stack.Push(5)
stack.Push(7)
stack.Push(4)
stack.Push(5)
```

**Извлечение**:

```go
stack.Pop() // 5
stack.Pop() // 7
stack.Pop() // 5
stack.Pop() // 4
```

**Объяснение**:

1. Первым извлекается элемент `5`, так как он самый частотный в стеке. При этом извлекается та `5`, которая была
   добавлена последней (по принципу стека)
2. Вторым извлекается элемент `7`, так как в стеке оба элемента `5` и `7` встречаются по 2 раза, но последняя `7` была
   добавлена позже, чем последняя `5` (`5`, которую мы добавили последним `Push`'ом мы извлекли первым вызовом `Pop`)
3. Третьим элементом извлекается элемент `5`, так как теперь это самый частотный эоемент в стеке
4. Четвертым извлекается элемент `4`, так как теперь в стеке все элементы имеют одинаковую частоту (каждый элемент
   встречается один раз), а элемент `4` был добавлен в стек последним

{% endtab %}
{% endtabs %}

---

## Решение

Реализуемая структура данных очень похожа на классическую задачу из мира практической разработки — `Priority Queue`. 
Однако в нашем случае в качестве приоритета выступает частота объекта, а алгоритмом извлечения является стек (`LIFO`), 
а не очередь (`FIFO`).

Подобно Priority Queue, одной из наиболее оптимальных структур данных для реализации будет Heap (куча).

#### Свойства кучи:

- Каждый узел имеет значение ключа больше/меньше (в зависимости от того, реализуем ли мы Max-кучу или Min-кучу) или равно значению ключей своих потомков.
- Куча является почти полным бинарным деревом. Это означает, что все уровни дерева полностью заполнены, кроме, возможно, последнего уровня, который заполняется слева направо.

#### Реализация кучи с помощью массива

Куча является древовидной структурой, которую можно эффективно реализовать с помощью массива, вычисляя индексы элементов следующим образом:

- Родительский элемент узла с индексом `i` находится по индексу `(i - 1) / 2`.
- Левый потомок узла с индексом `i` находится по индексу `2 * i + 1`.
- Правый потомок узла с индексом `i` находится по индексу `2 * i + 2`.

#### Как реализовать требуемую структуру данных

Для реализации Maximum Frequency Stack нам потребуются следующие структуры данных:
```go
type FreqStack struct {
	h        heap
	freqElem map[int]uint16
	curSeq   uint16
}

type heap []*Item

type Item struct {
	freq uint16
	seq  uint16
	val  int
}
```

`FreqStack` - основная структура, которую нам требуется реализовать. Она состоит из:
- `h` - куча, реализованная с помощью массива.
- `freqElem` - хеш-мапа, позволяющая определить частоту элемента на момент вставки.
- `curSeq` - текущий максимальный номер элемента (счетчик элементов).

У этой структуры нам необходимо реализвать методы `Push(val int)` и `Pop() int`.

### Push(val int)

Чтобы вставить элемент в кучу, нам необходимо определить ключ элемента (именно по ключам сравниваются элементы в куче).
В нашем случае ключ будет состоять из двух значений:
- Частота элемента в куче (можем быстро определять с помощью `freqElem`). При этом нам достаточно знать частоту элемента на момент вставки (не нужно обновлять уже вставленные элементы в кучу и изменять их положение).  
- Пордковый номер элемента в куче (для определение используется `curSeq`).

В кучу элемент изначально помещается в самый конец, после запускается рекурсивный процесс "просеивания вверх" элементов
кучи, который восстанавливает свойства кучи.

#### Просеивание вверх элементов в куче (upElement)

Просеивание вверх используется после вставки нового элемента в кучу. Элемент вставляется в конец массива (последний уровень дерева), и затем он перемещается вверх до тех пор, пока не будет восстановлено свойство кучи.

Алгоритм просеивания вверх:

1.	Найдите родительский элемент для последнего вставленного элемента. Поскольку мы реализовали кучу через массив, родительский элемент будет располагаться в массиве по индексу (i - 1) / 2.
2.	Сравните ключ нового элемента с его родительским узлом. В нашем случае ключ состоит из частоты и порядкового номера. Один элемент будет считаться больше другого, если его частота больше (при условии, что частоты у двух элементов разные), либо если его порядковый номер вставки больше.
3.	Если элемент больше родительского узла, поменяйте их местами.
4.	Повторяйте шаги 2 и 3, пока элемент не окажется на своем месте или не достигнет корня.

Таким образом, общая сложность вставки элемента в кучу определяется сложностью операции просеивания вверх, которая 
составляет `O(log n)`, так как в худшем случае нам придется пройтись вверх по бинарному дереву до корневого элемента.

### Pop() int

При извлечении элемента нам необходимо уменьшить частоту этого элемента, а также извлечь его из кучи.

Благодаря свойствам кучи необходимый для вставки элемент всегда будет находится в корне (в нашем случае в массиве под индексом `0`).
Аналогично с процессом добавления нового элемента, после извлечения корневого элемента из кучи, нам необходимо 
запустить процесс "просеивания вниз" элементов кучи, который восстанавливает свойства кучи.

Для этого:
1. Последний элемент кучи перемещается на первое место. Таким образом мы избегаем необходимость сдвигать все элементы 
кучи на один элемент влево.
2. Родительский элемент сравнивается с дочерними
3. Если элемент меньше одного из дочерних узлов, наибольший из дочерних элементов меняется местами с родительским.
4.	Повторяйте шаги 2 и 3, пока элемент не окажется на своем месте или не достигнет последнего уровня дерева.

Аналогично с процессом вставки, сложность процесса извлечения составляет `O(log n)`.

### Реализация

{% renderFile "_includes/components/solution.njk", taskName = "maximum_frequency_stack" %}

### Оценка сложности

**По времени**

Процесс вставки и извлечения элементов в куче составляет `O(log n)`, так как в обоих процессах нам необходимо запускать 
рекурсивный процесс "просеивания", который идет вверх п одереву от листового элемента (в случае просеивания вверх), либо 
вниз от корневого элемента до листового (в случае просеивания вниз).

**По памяти**

Доп память в данном случае составляет `O(n)`, так как мы используем дополнительную хеш-мапу `freqElem`.
Размер данной хеш-мапы в худшем случае равен количеству вставленных элементов.