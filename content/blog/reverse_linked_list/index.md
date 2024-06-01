---
layout: layouts/post.njk
title: Разворот связанного списка
date: 2023-12-04
complexity: easy
original_url: https://leetcode.com/problems/reverse-linked-list
preview_image: /images/previews/reverse_linked_list.webp
tags:
  - easy
  - linked list
  - recursion
---
---

## Описание задачи

Вам дан односвязный список. Переверните список и верните его.

Список представлен следующей структурой.

{% tabs %}

{% tab "GO" %}

```go
type ListNode struct {
    Val  int
    Next *ListNode
}
```

{% endtab %}

{% tab "Type Script" %}

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

{% tabs %}

{% tab "Пример №1" %}

![Связный список 1](/images/resources/linked_list_ex1.jpg)

{% endtab %}

{% tab "Пример №2" %}

![Связный список 2](/images/resources/linked_list_ex2.jpg)

{% endtab %}
{% endtabs %}

---

## Решение

Для реализации данной задачи нам всего лишь необходимо перебрать все узлы связанного списка и поменять местами указатели на следующий и предыдущий узлы.

### Реализация

{% renderFile "_includes/components/solution.njk", taskName = "reverse_linked_list" %}

### Оценка сложности

**По времени**

Сложность `O(n)`, так как мы итерируемся по всем узлам списка.

**По памяти**

Сложность `O(1)`, так как мы не выделяем дополнительную память.
