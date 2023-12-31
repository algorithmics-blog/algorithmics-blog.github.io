---
layout: post
title:  "Удаление дубликатов из отсортированного массива"
complexity: easy
original_url: https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/
tag: easy
---

## Оглавление

- [Описание задачи](#описание-задачи)
- [Ограничения](#ограничения)
- [Примеры](#примеры)
- [Решение](#решение)

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

{% tabs remove_duplicates_from_sorted_array %}

{% tab remove_duplicates_from_sorted_array Пример №1 %}
**Входные данные**: `[1, 1, 2]`

**Ответ**: `2`

В массиве два уникальных числа: 1 и 2.
{% endtab %}

{% tab remove_duplicates_from_sorted_array Пример №2 %}
**Входные данные**: `[0, 0, 1, 1, 1, 2, 2, 3, 3, 4]`

**Ответ**: `5`

В массиве пять уникальных чисел: 0, 1, 2, 3, 4.
{% endtab %}

{% endtabs %}

---

## Решение

По условиям задачи алгоритм не может выделять дополнительную память во время выполнения, поэтому мы не можем использовать дополнительные структуры даных, такие как `Маp` для подсчета уникальных значений.
Вместо этого мы сделаем указатель `pivot`, который будет отслиживать позицию последнего уникального элемента в измененном массиве.

- Устанавливаем значение `pivot = 0`.
- Запускаем цикл, который будет итерироваться по всем элементам в исходном массиве кроме последнего.

Так как массив отсортирован по возрастанию мы можем очень просто определять появление нового уникального элемента.
Если мы встречаем два соседних элемента, которые не равны друг другу, это означает, что они уникальны.

Если значение элемента под индексом `i` не равно значению элемента под индексом `i + 1`, то мы нашли новый уникальный элемент.
В этом случае нужно перезаписать значение элемента под индексом `pivot` значением по индексу `i` и увеличить `pivot` на 1.
Таки образом мы всегда переносим уникальные элементы по порядку в начало массива.

В итоге значение `pivot` становится равно количенству уникальных элемнетов, которые мы нашли.
Однако, он не учитывает последний элемент массива.
Например, в таком массиве `[1, 2, 3, 3]` значение `pivot` в конце будет равно `1`, так как мы не смещали индекс для цифры `3`.
Поэтому, в качестве ответа достаточно возвращать значение `pivot + 1`.

{% include code-example.md go_path="go/solution.go" ts_path="ts/solution.ts" %}

## Оценка сложности

`n` - количество элементов в массиве.

**По времени**

Сложность по времени `O(n)`, так как мы итерируемся по всем элементам массива.

**По памяти**

Сложность по памяти `O(1)`, так как мы не создаем дополнительных переменных.
