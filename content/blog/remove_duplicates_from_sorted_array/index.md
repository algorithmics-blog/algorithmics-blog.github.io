---
layout: layouts/post.njk
title: Удаление дубликатов из отсортированного массива
date: 2023-08-18
complexity: easy
original_url: https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/
preview_image: /images/previews/remove_duplicates_from_sorted_array.webp
tags:
  - easy
  - array
  - two pointers
---
---

## Описание задачи

Дан массив целых чисел `nums`, отсортированный в возрастающем порядке.
Напишите функцию `removeDuplicates` которая будет удалять дублирующиеся числа из массива `in-place`, то есть без выделения дополнительной памяти.
Относительный порядок элементов должен быть сохранен.

Измените массив `nums` так, чтобы первые `k` элементов в нем содержали уникальные элементы в том порядке, в котором они изначально присутствовали в `nums`.
Остальные элементы `nums` не важны, как и размер `nums`.

В качестве ответа функция должна возвращать целое число, равное количеству уникальных чисел в массиве.

---

## Ограничения

- В массиве может быть от 1 до 10<sup>4</sup> элементов
- В качестве значений могут быть числа в диапазоне от -100 до 100
- Массив отсортирован в возрастающем порядке

---

## Примеры

{% tabs %}
{% tab "Пример №1" %}

**Входные данные**: `[1, 1, 2]`

**Ответ**: `2`

В массиве два уникальных числа: 1 и 2.

{% endtab %}

{% tab "Пример №2" %}

**Входные данные**: `[0, 0, 1, 1, 1, 2, 2, 3, 3, 4]`

**Ответ**: `5`

В массиве пять уникальных чисел: 0, 1, 2, 3, 4.

{% endtab %}
{% endtabs %}

---

## Решение

По условиям задачи алгоритм не может выделять дополнительную память во время выполнения, поэтому мы не можем использовать дополнительные структуры даных, такие как `Маp` для подсчета уникальных значений.
Вместо этого мы сделаем указатель `pivot`, который будет отслеживать позицию для вставки нового уникального элемента в измененном массиве.

- Устанавливаем значение `pivot = 1`, так как нулевой элемент менять не нужно. Он уже на своем месте
- Запускаем цикл, который будет итерироваться по всем элементам в исходном массиве кроме первого.

Так как массив отсортирован по возрастанию мы можем очень просто определять появление нового уникального элемента.
Если мы встречаем два соседних элемента, которые не равны друг другу, это означает, что они уникальны.

Если значение элемента под индексом `i` не равно значению элемента под индексом `i - 1`, то мы нашли новый уникальный элемент.
В этом случае нужно перезаписать значение элемента под индексом `pivot` значением по индексу `i` и увеличить `pivot` на 1.
Таки образом мы всегда переносим уникальные элементы по порядку в начало массива.

В итоге значение `pivot` становится равно количеству уникальных элементов, которые мы нашли.

{% renderFile "_includes/components/solution.njk", taskName = "remove_duplicates_from_sorted_array" %}

## Оценка сложности

`n` - количество элементов в массиве.

**По времени**

Сложность по времени `O(n)`, так как мы итерируемся по всем элементам массива.

**По памяти**

Сложность по памяти `O(1)`, так как мы не создаем дополнительных переменных.