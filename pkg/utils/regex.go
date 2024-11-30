package utils

import "regexp"

func Scan(str, keyword string, i int) string {
	re := regexp.MustCompile(keyword)
	match := re.FindStringSubmatch(str)
	var result string
	if len(match) > i {
		result = match[i]
	}
	return result
}

func Scans(str, keyword string, i int) []string {
	re := regexp.MustCompile(keyword)
	matches := re.FindAllStringSubmatch(str, -1)

	var results []string
	for _, match := range matches {
		results = append(results, match[i])
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
