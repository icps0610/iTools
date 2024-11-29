package diff

import (
	"fmt"

	"tools/pkg/utils"
)

type Diff struct {
	Idx      int
	Content1 string
	Content2 string
}

func Run(content1, content2 []string) ([]Diff, []Diff, []Diff) {
	var diff, same, all []Diff

	max := utils.Max(len(content1), len(content2))
	for idx := 0; idx < max; idx++ {

		var element1, element2 string
		if len(content1) > idx {
			element1 = content1[idx]
		}
		if len(content2) > idx {
			element2 = content2[idx]
		}

		compare := Diff{idx + 1, element1, element2}

		if element1 == element2 {
			same = append(same, compare)
		} else {
			diff = append(diff, compare)
		}

		all = append(all, compare)
	}

	return diff, same, all
}

func Print(lists []Diff) {
	fmt.Println()
	for _, list := range lists {
		msg := fmt.Sprintf(`%v | %s | %s`, list.Idx, list.Content1, list.Content2)

		fmt.Println(msg)
	}
}
