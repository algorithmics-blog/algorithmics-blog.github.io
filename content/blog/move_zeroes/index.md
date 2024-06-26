---
layout: layouts/post.njk
title: Переместите нули
date: 2024-03-09
complexity: easy
original_url: https://leetcode.com/problems/move-zeroes/
preview_image: /images/previews/move_zeroes.webp
tags:
  - array
  - two pointers
  - easy
---
---

## Описание задачи

Вам дан целочисленный массив `nums`.
Переместите в нём все нули в конец, сохраняя относительный порядок ненулевых элементов.

Обратите внимание, что вы должны сделать это in-place, не копируя массив.

---

## Ограничения

- Длина массива от 1 до 10000 элементов
- Каждый элемент массива может принимать значение в диапазоне от -2<sup>31</sup> до 2<sup>31</sup> - 1

---

## Примеры

{% tabs %}

{% tab "Пример №1" %}

**Входные данные**: `nums = [0,1,0,3,12]`

**Ответ**: `[1,3,12,0,0]`

{% endtab %}

{% tab "Пример №2" %}

**Входные данные**: `nums = [0]`

**Ответ**: `[0]`

{% endtab %}
{% endtabs %}

---

## Решение

Решение задачи крайне простое.

Нам нужно завести два индекса.
Первый — обычный индекс `i`, который инкриминируется на каждой итерации.
Второй `firstZeroIndex` — индекс первого нулевого элемента в массиве, который изначально равен `0`.

Далее мы итерируемся по всем элементам массива и, как только встречаем ненулевой элемент, меняем его местами с нулем, который стоит под индексом `firstZeroIndex`.
Дополнительно увеличиваем `firstZeroIndex`, если произошла перестановка.

Таким образом все нули будут «всплывать» в конец массива за один проход.

### Реализация

{% renderFile "_includes/components/solution.njk", taskName = "move_zeroes" %}

### Оценка сложности

**По времени**

Сложность `O(n)`, так как мы итерируемся по всем элементам массива.

**По памяти**

Сложность `O(1)`, так как мы не выделяем дополнительную память, зависящую от длины массива.
