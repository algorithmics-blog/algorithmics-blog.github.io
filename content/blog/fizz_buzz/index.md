---
layout: layouts/post.njk
title: Fizz Buzz
date: 2023-12-01
complexity: easy
original_url: https://leetcode.com/problems/fizz-buzz/
preview_image: /images/previews/fizz_buzz.webp
tags:
  - math
  - string
  - easy
---
---

## Описание задачи

Дано целое число `n`. 
Напишите функцию, которая принимает в качестве параметра число `n` и возвращает массив строк `answer`, который формируется по следующим правилам:

- `answer[i] == "FizzBuzz"`, если `i` делится одновременно на 3 и 5;
- `answer[i] == "Fizz"`, если `i` делится только на 3;
- `answer[i] == "Buzz"`, если `i` делится только на 5;
- `answer[i] == i`, во всех остальных случаях.

---

## Ограничения

Значение `n` находится в диапазоне от 1 до 10<sup>4</sup>.

---

## Примеры

{% tabs %}

{% tab "Пример №1" %}

**Входные данные**: `3`

**Ответ**: `["1","2","Fizz"]`

{% endtab %}

{% tab "Пример №2" %}

**Входные данные**: `5`

**Ответ**: `["1","2","Fizz","4","Buzz"]`

{% endtab %}

{% tab "Пример №3" %}

**Входные данные**: `14`

**Ответ**: `["1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz","11","Fizz","13","14","FizzBuzz"]`

{% endtab %}
{% endtabs %}


---

## Решение

Для решения задачи нужно запустить цикл от 1 до `n` и на каждой итерации проверять 4 взаимоисключающих условия.
Если условие выполняется, то в результирующий массив добавляется соответсвующая строка.

### Реализация

{% renderFile "_includes/components/solution.njk", taskName = "fizz_buzz" %}

### Оценка сложности

**По времени**

Сложность по времени линейная — `O(n)`, так как мы итерируемся от 1 до `n`.

**По памяти**

Сложность по памяти константная — `O(1)`, так как мы не создаем дополнительных переменных.
