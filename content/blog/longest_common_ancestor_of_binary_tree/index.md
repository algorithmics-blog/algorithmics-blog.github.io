---
layout: layouts/post.njk
title: Самый низкий общий предок двоичного дерева
date: 2024-02-20
complexity: medium
original_url: https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree
preview_image: /images/previews/longest_common_ancestor_of_binary_tree.webp
tags:
  - medium
  - tree
  - binary tree
---
---

## Описание задачи

Вам дано двоичное дерево.
Найдите наименьшего общего предка (LCA) двух заданных узлов в дереве.

Согласно [определению LCA в Википедии](https://en.wikipedia.org/wiki/Lowest_common_ancestor): «Наименьший общий предок
определяется между двумя узлами `p` и `q` как самый нижний узел в дереве `T`, который имеет как `p`, так и `q` в качестве
потомков (где мы позволяем узлу быть потомком самого себя).»

---

## Ограничения

- Количество узлов в дереве находится в диапазоне от 2 до 10<sup>5</sup>.
- Значения каждого узла в дереве находятся в диапазоне от -10<sup>9</sup> до 10<sup>9</sup>.
- Значения всех узлов в дереве уникальны.
- `p != q`.
- `p` и `q` всегда существуют в дереве.

---

## Примеры

{% tabs %}
{% tab "Пример №1" %}

![Бинарное дерево](/images/resources/binarytree.png)

**Входные параметры**: дерево выше, `p = 5`, `q = 1`.

**Ответ**: `3`

**Объяснение**: LCA узлов 5 и 1 равен 3.

{% endtab %}

{% tab "Пример №2" %}

![Бинарное дерево](/images/resources/binarytree.png)

**Входные параметры**: дерево выше, `p = 5`, `q = 4`.

**Ответ**: `5`

**Объяснение**: LCA узлов 5 и 4 равен 5, поскольку узел может быть потомком самого себя согласно определению LCA.

{% endtab %}
{% endtabs %}

---

## Решение

Решение этой задачи достаточно интуитивно.
Сначала мы пройдем по дереву вглубь.
В тот момент, когда встретится любой из узлов `p` или `q`, вернем логический флаг.
Флаг помогает определить, нашли ли мы нужные узлы на каком-либо из путей.
Тогда наименьшим общим предком будет узел, для которого обе рекурсии поддерева возвращают флаг `true`.
Это также может быть узел, который сам является одним из узлов `p` или `q` и для которого одна из рекурсий поддерева возвращает флаг `true`. 

### Реализация

Для реализации будем использовать рекурсивный подход с замыканием для хранения переменной `lca`.

{% renderFile "_includes/components/solution.njk", taskName = "longest_common_ancestor_of_binary_tree" %}

### Оценка сложности

**По времени**

`O(n)` так как в худшем случае нам нужно посетить все `n` узлов в дереве.

**По памяти**

`O(n)` так как максимальный объем пространства, используемого стеком рекурсии, будет равен `n`.