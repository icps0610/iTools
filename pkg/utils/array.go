package utils

func GetRange(min, max int) []int {
	var list []int
	for idx := min; idx <= max; idx++ {
		list = append(list, idx)
	}
	return list
}
