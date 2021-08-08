package utilservice

import (
	utilrepository "github.com/axel526/jikkosoft/src/data/repositories/util.repository"
	e "github.com/axel526/jikkosoft/src/entity"
)

func Sorted(numbers []int) e.Sorted {
	var sorted e.Sorted

	sorted.Unsorted = numbers
	onlyNumbers, repeatNumbers := utilrepository.SeparateRepeatedNumbers(sorted.Unsorted)
	onlyNumbers = utilrepository.OrderNumbers(onlyNumbers)
	sorted.Sorted = append(onlyNumbers, repeatNumbers...)

	return sorted
}
