---
layout: post
title:  "Разворот связанного списка"
complexity: easy
original_url: https://leetcode.com/problems/reverse-linked-list
tag: easy
---

## Оглавление

- [Описание задачи](#описание-задачи)
- [Ограничения](#ограничения)
- [Примеры](#примеры)
- [Решение](#решение)

---

## Описание задачи

Вам дан односвязный список. Переверните список и верните его.

Список представлен следующей структурой.

{% tabs reverse_linked_list_node %}

{% tab reverse_linked_list_node GO %}
```go
type ListNode struct {
    Val  int
    Next *ListNode
}
```
{% endtab %}

{% tab reverse_linked_list_node Type Script %}
```typescript
class ListNode {
    val: number
    next: ListNode | null
}
```
{% endtab %}
{% endtabs %}

---

## Ограничения

- Количество узлов в связанном списке находится в диапазоне от 1 до 5000
- Значение каждого узла находится в диапазоне от -5000 до 5000

---

## Примеры

{% tabs move_zeroes %}

{% tab move_zeroes Пример №1 %}

![Связный список 1](/assets/images/linked_list_ex1.jpg)

{% endtab %}

{% tab move_zeroes Пример №2 %}

![Связный список 2](/assets/images/linked_list_ex2.jpg)

{% endtab %}
{% endtabs %}

---

## Решение

Для реализации данной задачи нам всего лишь необходимо перебрать все узлы связанного списка и поменять местами указатели на следующий и предыдущий узлы.

### Реализация

{% include code-example.md go_path="go/solution.go" ts_path="ts/solution.ts" %}

### Оценка сложности

**По времени**

Сложность `O(n)`, так как мы итерируемся по всем узлам списка.

**По памяти**

Сложность `O(1)`, так как мы не выделяем дополнительную память.
