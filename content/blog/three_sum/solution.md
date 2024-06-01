---
layout: layouts/post.njk
title: Сумма трех чисел в массиве
date: 2023-09-06
complexity: medium
original_url: https://leetcode.com/problems/3sum
preview_image: /images/previews/three_sum.webp
tags:
  - medium
  - array
  - two pointers
---
---

## Описание задачи

Дан массив целых чисел `nums`.
Напишите функцию, чтобы найти в массиве все уникальные тройки чисел, сумма которых равна нулю.
В одной тройке `[nums[i], nums[j], nums[k]]` один и тот же элемент не может повторяться дважды, то есть `i != j, i != k, j != k`.

Обратите внимание, что в ответе не должно быть повторяющихся троек. Порядок троек и чисел в тройках не имеет значения.

---

## Ограничения

- В массиве может быть от 3 до 3000 значений
- В качестве значений могут быть числа в диапазоне от -10<sup>5</sup> до 10<sup>5</sup>

---

## Примеры

{% tabs %}

{% tab "Пример №1" %}

**Входные данные**: `[-1, 0, 1, 2, -1, -4]`

**Ответ**: `[[-1, -1, 2], [-1, 0, 1]]`

{% endtab %}

{% tab "Пример №2" %}

**Входные данные**: `[0, 0, 0]`

**Ответ**: `[[0, 0, 0]]`

{% endtab %}

{% tab "Пример №3" %}

**Входные данные**: `[0, 1, 1]`

**Ответ**: `[]`

{% endtab %}

{% endtabs %}

---

## Решение через Hash Set

Это решение требует предварительной сортировки массива в возрастающем порядке, поэтому первым делом выполняем сортировку.

Далее мы можем заметить, что задача очень похожа на ее более простой аналог — [Сумма двух чисел в массиве](/blog/two_sum/solution/).
Разница только в том, что нам надо найти все тройки, а их сумма всегда равно нулю.
Мы можем воспользоваться подходом из этой задачи и решить ее через дополнительную структуру `HashSet` или ее аналоги.

Посмотрим на такой пример ввода `[-1, 0, 1, 2, -1, -4]`.

После сортировки мы получим следующий массив `[-4, -1, -1, 0, 1, 2]`.

Теперь запускаем цикл по всем элементам массива.

Далее для каждого `i`-го элемента мы будем запускать отдельный цикл для поиска недостающей пары чисел из тройки.
Так как мы всегда двигаемся слева направо, то мы можем запускать второй цикл начиная от `i + 1` и до конца.

Первое, что нам нужно учесть — это то, что в наших ответах не должно быть одинаковых троек.
Если мы посмотрим еще раз на пример входящих данных, то мы обнаружим что у нас есть два одинаковых возможных решения `[-1, 0, 1]`.
В массиве есть две возможные тройки со значениями под индексами `1, 3, 4` и `2, 3, 4`.
Однако, это легко решить, так как у нас отсортированный массив. Это значит, что одинаковые значения идут подряд и мы просто можем проверять первое и пропускать все остальные.

Добавим проверку в самое начало первого цикла. Если мы перебираем не первый элемент в массиве и текущий элемент равен предыдущему,
то мы просто его пропускаем. И дополнительно создадим `Set` для отслеживания уже проверенных чисел.

Теперь у нас есть такая заготовка.

```typescript
export const threeSumHashSet = (nums: number[]): number[][] => {
    // Создаем массив для результатов
    const result: number[][] = []

    // Предварительно сортируем массив по возрастанию
    nums.sort()

    for (let i = 0; i < nums.length; i++) {
        const firstNum = nums[i]

        // Не проверяем текущее число, если оно такое же, как и предыдущее, потому для него мы получим такой же результат.
        if (i != 0 && nums[i - 1] == firstNum) {
            continue
        }

        const used = new Set<number>()

        for (let j = i + 1; j < nums.length; j++) {
            const secondNum = nums[j]

            // Имплементация
        }
    }

    return result
}
```

Осталось решить, что делать внутри второго цикла. А работа внутри него в целом сводится к задаче [Сумма двух чисел в массиве](/blog/two_sum/solution/).
Нам нужно вычислить третье число, которого нам не хватает, чтобы получить в сумме ноль и проверить, встречалось ли уже такое число в нашем массиве.
Для проверки мы как раз и используем наш `Set`. 

- Если третье число есть в `used`, то мы добавляем в наш результат массив из всех трех полученных чисел.
- Если третьего числа нет в `used`, то мы просто записываем в него второе число и идем дальше.

Осталось учесть еще один нюанс. Во время перебора второго массива мы тоже можем получать одинаковые значения и тем самым создавать дубликаты.
Поэтому, после добавления найденной тройки нам надо увеличивать `j` до тех пор, пока значение под индексом `j` равно значению под индексом `j-1`.

В результате мы получаем следующие решения задачи.

{% renderFile "_includes/components/solution.njk", taskName = "three_sum", fileName="solution_hash_set" %}

### Оценка сложности

`n` - количество элементов в массиве

**По времени** 

<code>O(n<sup>2</sup>)</code>, так как мы запускаем цикл в цикле. Дополнительную сложность добавляет сортировка массива, ее сложность равна `O(n * log n)`.
Итоговая сложность равна `O(n * log n)` + <code>O(n<sup>2</sup>)</code>, что асимптотически эквивалентно <code>O(n<sup>2</sup>)</code>.

**По памяти**

`O(n)`, так как мы добавляем в сет `n` элементов.

---

## Решение через два указателя

Это решение, так же как и предыдущее, требует предварительной сортировки массива.
Поэтому первым делом необходимо отсортировать массив в порядке возрастания.

Теперь мы можем применить подход с двумя указателями, как в задаче [Сумма двух чисел в отсортированном массиве](/blog/two_sum_sorted_array/solution/). 
Разница только в том, что нам надо найти все тройки, а их сумма всегда равно нулю.

Посмотрим на такой пример ввода `[-1, 0, 1, 2, -1, -4]`.

После сортировки мы получим следующий массив `[-4, -1, -1, 0, 1, 2]`.

Теперь запускаем цикл по всем элементам массива.

Далее для каждого `i`-го элемента мы будем запускать отдельный цикл для поиска недостающей пары чисел из тройки.
Так как мы всегда двигаемся слева направо, то мы можем запускать второй цикл начиная от `i + 1` и до конца.

Как и в предыдущем решении нам надо обеспечить отсутствие дубликатов в ответе, поэтому мы будем проверять одинаковые значения только один раз, а повторения пропускать.

Теперь у нас есть такая заготовка.

```typescript
export const threeSumTwoPointers = (nums: number[]): number[][] => {
    const result: number[][] = []

    nums.sort()

    for (let i = 0; i < nums.length; i++) {
        const num = nums[i]

        if (i != 0 && nums[i-1] == num) {
            continue
        }

        // Имплементация
    }

    return result
}
```

Теперь мы можем полностью скопировать решение из задачи [Сумма двух чисел в отсортированном массиве](/blog/two_sum_sorted_array/solution/) и немного его модифицировать.

1. Заводим два указателя, которые будут хранить индексы. Значение `left` делаем равным `i + 1`, значение `right` равным
   индексу последнего элемента.
2. Запускаем цикл, который прервется, если `left` станет больше или равен `right`, то есть когда индексы сойдутся.
3. На каждой итерации высчитываем сумму элементов под индексами `left`, `right`, и `i`-го элемента массива и проверяем ее на равенство нулю.

- Если сумма больше нуля, значит мы сложили слишком большие числа и нам надо взять более маленькие.
  Уменьшить сумму мы можем взяв число, которое стоит левее от текущего под индексом `right`.
  Уменьшаем `right` на 1.
- Если сумма меньше нуля, значит мы сложили слишком маленькие числа и нам надо взять более большие.
  Увеличить сумму мы можем взяв число, которое стоит правее от текущего под индексом `left`.
  Увеличиваем `left` на 1.
- Если сумма равна нулю, значит мы нашли искомые числа. В таком случае добавляем найденную тройку `[nums[i], nums[left], nums[right]]` в коллекцию ответов.
Здесь же сразу уменьшаем `right` и увеличиваем `left` на единицу, чтобы найти следующую возможную комбинацию.

Осталось учесть один нюанс, как и в предыдущем решении. Во время второго перебора мы можем получать одинаковые значения и тем самым создавать дубликаты.
Поэтому, после добавления найденной тройки нам надо увеличивать указатель `left` до тех пор, пока значение под индексом `left` равно значению под индексом `left - 1`.

{% renderFile "_includes/components/solution.njk", taskName = "three_sum", fileName="solution_two_pointers" %}

### Оценка сложности

`n` - количество элементов в массиве

**По времени**

<code>O(n<sup>2</sup>)</code>, так как мы запускаем цикл в цикле. Дополнительную сложность добавляет сортировка массива, ее сложность равна `O(n * log n)`.
Итоговая сложность равна `O(n * log n)` + <code>O(n<sup>2</sup>)</code>, что асимптотически эквивалентно <code>O(n<sup>2</sup>)</code>.

**По памяти**

Варьируется от `O(log n)` до `O(n)` в зависимости от реализации алгоритма сортировки.
На хранение указателей мы тратим константную память `O(1)`, поэтому финальная сложность по памяти определяется сложностью алгоритма сортировки.

---

## Решение без сортировки

Для решения задачи этим способом не требуется предварительная сортировка входного массива чисел.
Для решения можно немного видоизменить [реализацию через Hash Set](#reshenie-cherez-hash-set).

Необходимо так же проитерироваться по всем элементам массива и на каждой итерации запустить отдельный цикл от `i + 1`.
Для того чтобы не проверять повторяющиеся числа, мы будем записывать каждое в структуру Hash Set `dups` и в начале итерации проверять, если уже такой элемент.
Если в Hash Set уже есть такое число, то мы просто пропускаем итерацию.

Также необходимо определить Hash Map `seen`, в которую мы будем записывать уже проверенные пары двоек `firstNum` и `secondNum`.
При таком подходе мы можем вычислять на каждой итерации второго цикла недостающее число `thirdNum` и проверять есть ли уже такая пара в `seen`.
Если такая пара есть, то мы нашли искомую тройку.

Теперь нам нужно избежать добавления одинаковых троек в ответ.
Для этого можно воспользоваться хитрым приемом, но он работает не во всех языках. 

В таких языках, как `Python` и `Java` мы можем создать массив (лист) из трех элементов, отсортировать их в любом порядке и добавить в еще один Has Set `uniq`.
Добавление отсортированных троек обеспечит автоматическое избавление от дублей благодаря механике Hash Set.
В конце просто нужно конвертировать Hash Set в массив и вернуть в качестве ответа.

Однако, в таких языках как, например, `Go`, `JavaScript` и `TypeScript` провернуть такое не получится.
В `Go` нельзя использовать слайсы в качестве ключей, потому что это не сравниваемые типы.
В `JavaScript` и `TypeScript` массивы использовать можно, но они передаются по ссылке, поэтому в Hash Set можно спокойно добавить одинаковые массивы, так как у них разные ссылки.

В этих языках в качестве ключа придется использовать строку вида `1.2.3`, где через точку записаны три числа из тройки в отсортированном порядке.
Таким образом мы создадим уникальный ключ и избежим дубликатов при добавлении.
Однако, при возврате результата нам придется обратно разобрать каждую строку на отдельные числа с приведением типов.

{% renderFile "_includes/components/solution.njk", taskName = "three_sum", fileName="solution_no_sort" %}

### Оценка сложности

`n` - количество элементов в массиве

**По времени**

<code>O(n<sup>2</sup>)</code> так как мы запускаем цикл в цикле.

Хотя асимптотическая сложность такая же, как в предыдущих решениях, этот алгоритм заметно медленнее.
Поиск в хеш-наборе, хотя и требует постоянного времени, является дорогостоящим по сравнению с прямым доступом к памяти.

**По памяти**

`O(n)` так как мы заводим дополнительные структуры для хранения чисел, которые зависят от длины входного массива.
