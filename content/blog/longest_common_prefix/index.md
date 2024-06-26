---
layout: layouts/post.njk
title: Поиск самого длинного общего префикса в строках
date: 2023-09-15
complexity: easy
original_url: https://leetcode.com/problems/longest-common-prefix/
preview_image: /images/previews/longest_common_prefix.webp
tags:
  - string
  - easy
---
---

## Описание задачи

Напишите функцию для поиска самой длинной строки общего префикса среди массива строк.
Если общего префикса нет, верните пустую строку `""`.

Входные данные: массив строк.

Выходные данные: строка, являющаяся самым длинным общим префиксом для всех строк массива.

---

## Ограничения

- В массиве строк может быть от 1 до 200 элементов
- Длина каждой строки от 0 до 200 символов
- Каждая строка состоит только из строчных английских букв

---

## Примеры

{% tabs %}
{% tab "Пример №1" %}

**Входные данные**: `["flower","flow","flight"]`

**Ответ**: `"fl"`

{% endtab %}

{% tab "Пример №2" %}

**Входные данные**: `["dog","racecar","car"]`

**Ответ**: `""`

{% endtab %}
{% endtabs %}

---

## Решение

Для того чтобы найти общий префикс у нескольких строк, нужно последовательно от 0-го символа
сравнивать соответствующие по индексу символы в каждой строке.
Если символы совпадают, значит текущий символ входит в общий префикс.
Как только мы находим первый несовпадающий символ или заканчивается одна из строк, значит общий префикс закончился.

Итоговый алгоритм:

- Выбираем одну из строк массива по которой будем итерироваться. Для удобства можно выбрать первую.
- Итерируемся по выбранной строке посимвольно. Для каждого символа запускаем новый цикл, в котором проверяем,
  что во всех строках на соответствующем индексе стоит один и тот же символ.
- Если символ по индексу во всех строках одинаковый, то добавляем его к строке-результату
- Прекращаем итерации и возвращаем префикс в виде строки как только встречается первый несовпадающий символ или если превышается длина любой из строк.

{% renderFile "_includes/components/solution.njk", taskName = "longest_common_prefix" %}

### Оценка сложности

**По времени**

Сложность `O(m * n)`, где `n` — количество строк в массиве, `m` — длина самой короткой строки.
Мы итерируемся по строке и сравниваем соответствующие текущему индексу символы в каждой из строк.

**По памяти**

Сложность `O(1)`, то есть константная, потому что мы выделяем память только для хранения переменных массива и результирующей строки.
