package statistics

// Avg возвращает среднее значение массива чисел.
func Avg(nums []float64) float64 {
	var sum float64
	for _, n := range nums {
		sum += n
	}
	return sum / float64(len(nums))
}

// MaxN возвращает наибольшее число из массива чисел
func MaxN(nums []float64) float64 {
	var maxNum float64
	for i := range nums {
		if i == 0 {
			continue
		}
		if nums[i-1] < nums[i] {
			maxNum = nums[i]
		}
	}
	return maxNum
}
