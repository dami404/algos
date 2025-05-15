package main

/*
	Вам дан массив целых чисел digits, где каждый элемент — цифра. Массив может содержать дубликаты.

	Вам необходимо найти все уникальные целые числа , которые удовлетворяют заданным требованиям:
	- Целое число состоит из конкатенации трех элементов digits в любом произвольном порядке.
	- Целое число не имеет начальных нулей.
	- Целое число четное.
	Например, если даны digits, [1, 2, 3] целые числа 132 и 312 следуют требованиям.

	Верните отсортированный массив уникальных целых чисел.
*/

/*
	Затраты времени: O(k*10^k), где k- количество цифр в проверяемом числе
	Затраты памяти: О(1)
*/

func isInArray(arr, num map[int]int) bool {
	for digit, count := range num {
		if count > arr[digit] {
			return false
		}
	}
	return true
}

func countDigits(n int) map[int]int {
	digitsNumber := make(map[int]int)
	for n > 0 {
		digitsNumber[n%10]++
		n /= 10
	}
	return digitsNumber
}

func findEvenNumbers(digits []int) []int {
	result := []int{}
	digitsNumber := make(map[int]int)
	for _, digit := range digits {
		digitsNumber[digit]++
	}
	for i := 100; i <= 999; i += 2 {
		if isInArray(digitsNumber, countDigits(i)) {
			result = append(result, i)
		}
	}
	return result
}
