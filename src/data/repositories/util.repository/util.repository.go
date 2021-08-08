package utilrepository

func OrderNumbers(numbers []int) []int {
	for i := range numbers {
		for j := range numbers {
			temp := 0

			if numbers[i] < numbers[j] {
				temp = numbers[i]
				numbers[i] = numbers[j]
				numbers[j] = temp
			}

		}
	}

	return numbers
}

func SeparateRepeatedNumbers(numbers []int) ([]int, []int) {
	var onlyNumbers, repeatNumbers []int

	for index, item := range numbers {
		if index == 0 {
			onlyNumbers = append(onlyNumbers, item)
			continue
		}

		exist := false

		for _, onlyNumber := range onlyNumbers {
			if onlyNumber == item {
				exist = true
				break
			}
		}

		if !exist {
			onlyNumbers = append(onlyNumbers, item)
		} else {
			repeatNumbers = append(repeatNumbers, item)
		}

	}

	return onlyNumbers, repeatNumbers
}
