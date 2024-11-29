package utils

func Max(num ...int) int {
	var max = num[0]
	for _, num := range num {
		if num > max {
			max = num
		}
	}
	return max
}
