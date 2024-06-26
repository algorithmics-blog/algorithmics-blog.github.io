---
layout: layouts/post.njk
title: Дневная температура
date: 2023-12-22
complexity: medium
original_url: https://leetcode.com/problems/daily-temperatures/
preview_image: /images/previews/daily_temperatures.webp
tags:
  - medium
  - array
  - stack
---
---

## Описание задачи

Напишите функцию, которая будет рассчитывать количество дней между текущим и днем с большей температурой.

**Входные данные**

Массив целых неотрицательных чисел. Каждое число в массиве, это температура в соответствующий день.

**Выходные данные**

Массив целых неотрицательных чисел. Каждое число в массиве, это количество дней от `i`'го дня до дня с большей температурой.

---

## Ограничения

- Количество дней (длина входного массива) находится в диапазоне от 1 до 100000
- Температура для каждого дня находится в диапазоне от 30 до 100

---

## Примеры

{% tabs %}

{% tab "Пример №1" %}

**Входные данные**: `[73,74,75,71,69,72,76,73]`

**Ответ**: `[1,1,4,2,1,1,0,0]`

{% endtab %}

{% tab "Пример №2" %}

**Входные данные**: `[30,40,50,60]`

**Ответ**: `[1,1,1,0]`

{% endtab %}

{% tab "Пример №3" %}

**Входные данные**: `[30,60,90]`

**Ответ**: `[1,1,0]`

{% endtab %}
{% endtabs %}

---

## Решение через стек

Задачу можно решить за линейное время с использованием структуры данных стек.

Решение очень похоже на предыдущие задачи: [столкновение астероидов](/blog/asteroid_collision) и [валидация скобочной последовательности](/blog/valid_parentheses/solution)

Для решения задачи, нам необходимо завести результирующий массив, стек (можно реализовать стек через массив) и в цикле перебрать элементы исходного массива:
- Результирующий массив полностью заполняем нулями (размер результирующего массива равен размеру исходного)
- Если стек не пустой, запускаем внутренний цикл, доставая элементы (индексы) из стека (для удобства в тексте обозначим индекс, который мы достали из верхушки стека как `i`, а индекс текущего дня как `j`):
  - Если температура `j`'го дня меньше температуры `i`'го дня, возвращаем индекс `i` в стек и добавляем в стек сверху индекс `j`, выходим из внутреннего цикла и переходим к следующему дню.
  - Если температура `j`'го дня больше, `i`'ый элемент результирующего массива приравниваем `j - i` (`j` - номер текущего дня по порядку, `i` - номер дня из стека. То есть, разница этих значений и будет количеством дней между двумя днями). Достаем следующий элемент из стека.
- Если стек пустой, добавляем `j` индекс элемента в стек.

{% renderFile "_includes/components/solution.njk", taskName = "daily_temperatures" %}

### Оценка сложности

**По времени**

Сложность `O(n)`, где `n` — длина строки.

В худшем случае нам придется положить `n-1` элементов в стек, а после извлечь все элементы. То есть, максимальная сложность равна `2n`, что с точки зрения оценки алгоритма эквивалентно `O(n)`.

**По памяти**

Сложность `O(n)`, так как нам пришлось создать стек и результирующий массив для хранения данных.
