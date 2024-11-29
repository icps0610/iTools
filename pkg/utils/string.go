package utils

import (
	"fmt"
	"regexp"
	"strconv"
)

func To_i(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func To_f(s string) float64 {
	fn, _ := strconv.ParseFloat(s, 64)
	return fn
}

func To_s(i int) string {
	return fmt.Sprintf(`%v`, i)
}

func Scan(str, keyword string, i int) string {
	re := regexp.MustCompile(keyword)
	match := re.FindStringSubmatch(str)
	var result string
	if len(match) > i {
		result = match[i]
	}
	return result
}

func Scans(str, keyword string) []string {
	re := regexp.MustCompile(keyword)
	matches := re.FindAllStringSubmatch(str, -1)

	var results []string
	for _, match := range matches {
		results = append(results, match[0])
	}
	return results
}

func ScanGroups(str, keyword string) []string {
	re := regexp.MustCompile(keyword)
	return re.FindStringSubmatch(str)
}

func Match(str, keyword string) bool {
	match, _ := regexp.MatchString(keyword, str)
	return match
}
