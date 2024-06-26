---
layout: layouts/post.njk
title: Максимальный средний подмассив
date: 2024-05-10
complexity: easy
original_url: https://leetcode.com/problems/maximum-average-subarray-i/
preview_image: /images/previews/maximum_average_subarray_i.webp
tags:
  - array
  - sliding window
  - easy
---
---

## Описание задачи

Вам дан целочисленный массив `nums`, состоящий из `n` элементов, и целое число `k`.
Найдите непрерывный подмассив длиной `k`, имеющий максимальное среднее значение, и верните это значение.

Принимается любой ответ с ошибкой расчета менее 10<sup>-5</sup>.

---

## Ограничения

- `k` гарантированно меньше или равно длине массива `nums`
- Количество элементов в массиве может быть в диапазоне от 1 до 10<sup>5</sup>
- Каждый элемент массива может принимать значение в диапазоне от -10<sup>4</sup> до 10<sup>4</sup>

---

## Примеры

{% tabs %}

{% tab "Пример №1" %}

**Входные данные**: `nums = [1,12,-5,-6,50,3], k = 4`

**Ответ**: `12.75`

{% endtab %}

{% tab "Пример №2" %}

**Входные данные**: `nums = [5], k = 1`

**Ответ**: `5`

{% endtab %}
{% endtabs %}

---

## Решение

Для решения задачи мы будем использовать метод скользящего окна.

Начнем мы с того, что посчитаем сумму первых `k` элементов массива таким образом, как бы предзаполнив значение скользящего окна длиною `k`.
После этого создадим переменную `maxSum`, которая будет хранить максимальное значение суммы подмассива и присвоим ей значение суммы первых `k` элементов.

Далее мы будем итерироваться по массиву, начиная с элемента `k` и до конца массива.
На каждой итерации мы будем вычитать из суммы текущего окна значение элемента, который выходит за пределы окна слева, и прибавлять значение элемента, который входит в окно справа.
После этого мы будем сравнивать текущее значение суммы окна с максимальным значением и обновлять `maxSum`, если текущее значение больше.

Таким образом, после завершения всех итераций по массиву, в переменной `maxSum` будет храниться максимальное значение суммы подмассива длиною `k`.
В самом конце только остается поделить эту сумму на `k` таким образом получив среднее значение подмассива.

### Реализация

{% renderFile "_includes/components/solution.njk", taskName = "maximum_average_subarray_i" %}

### Оценка сложности

**По времени**

Сложность `O(n)`, так как мы итерируемся по всем элементам массива.

**По памяти**

Сложность `O(1)`, так как мы не выделяем дополнительную память, зависящую от длины массива.
