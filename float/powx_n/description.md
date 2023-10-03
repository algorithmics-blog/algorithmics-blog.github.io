# Имплементация функции возведения в степень

[leetcode](https://leetcode.com/problems/powx-n/description/)

Сложность: Средняя

## Оглавление

- [Описание](#description)
- [Брутфорс решение](#solution_1)
- [Оптимальное решение](#solution_2)

---

## <a name="description"></a>Описание

Напишите функцию, которая будет возводить число x (float64) в степень n.

Входные данные: x (float64), n (int)

Выходные данные: float64 (результат возведения числа x в степень n)

### Ограничения

- -100.0 < x < 100.0
- -2<sup>31</sup> <= n <= 2<sup>31</sup>-1
- -10<sup>4</sup> <= x<sup>n</sup> <= 10<sup>4</sup>
- При x = 0, n - неотрицательные

### Примеры

#### Пример 1

```
Вход: x = 2.00000, n = 10
```

```
Ответ: 1024.00000
```

#### Пример 2

```
Вход: x = 2.10000, n = 3
```

```
Ответ: 9.26100
```

#### Пример 3

```
Вход: 2.00000, n = -2
```

```
Ответ: 0.25000
```

#### Каких примеров не хватает в описании (корнер кейсы)
- 1 в любой степени
- -1 в четной степени
- -1 в нечетной стпени


## <a name="solution_1"></a> Брутфорс решение

Так как возведение в степень — результат умножения числа `x` само на себя `n` раз, брутфорс решение можно написать используя обычный цикл (пока предположим, что `n` всегда > 0):
```go
res := x

for ; n > 1; n-- {
    res *= x
}

return res
```

Для того, чтобы учесть кейсы с отрицательным `n`, вспомним свойство степеней.

$$ x^{-n} = {1 \over x^n} $$

Таким образом, вынесем возведение в положительную степень в отдельную функцию:
```go
func powAbsNBruteforce(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if n == 1 {
		return x
	}

	res := x

	for ; n > 1; n-- {
		res *= x
	}

	return res
}
```
А в основной функции добавим проверку на отрицательность `n`:
```go
if n < 0 {
    return 1 / powAbsNBruteforce(x, -1*n)
}

return powAbsNBruteforce(x, n)
```

### Оценка сложности

#### Сложность по времени

Для вычисления степени нам нужно произвести n умножений, т.е. сложность -  **O(n)**

#### Сложность по памяти

`O(1)` — дополнительная память константна.

## <a name="solution_1"></a> Оптимальное решение

Мы можем сократить кол-во умножений, если вспомним одно из свойств степеней.

$$ x^n = {(x * x)^{n \over 2}} $$

Например:
<code>2<sup>8</sup>=4<sup>4</sup>=16<sup>2</sup>=256</code>.

Таким образом, мы можем переписать нашу функцию на рекурсивную, не забывая отдельно обработать кейс с нечетной степенью:
```go
func powAbsN(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if n == 1 {
		return x
	}

	if n%2 != 0 {
		return powAbsN(x*x, n/2) * x
	}

	return powAbsN(x*x, n/2)
}
```

### Оценка сложности

#### Сложность по времени

В этом подходе на каждой итерации рекурсии мы возводим число в квадрат, вдвое снижая количество требуемых умножений.
То есть итоговая сложность — `O(log(n))`

#### Сложность по памяти

`O(1)` — дополнительная память константна.