package diff

import (
	"fmt"

	"strings"
)

func OutputHtml(path1, path2 string, difference, same, all []Diff) string {
	htmlText := fmt.Sprintf(`<!DOCTYPE html><html><head><meta charset="utf-8"><meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no"><link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-QWTKZyjpPEjISv5WaRU9OFeRpok6YctnYmDr5pNlyT2bRjXh0JMhjY6hW+ALEwIH" crossorigin="anonymous"><title>Diff</title></head><body><div class="container-fluid"><div class="row"><div class="col-1"></div><div class="col-10"><br>`)

	htmlText += tableText(fmt.Sprintf(`%s vs %s`, path1, path2), `primary`, all)
	htmlText += tableText(`same`, `info`, same)
	htmlText += tableText(`difference`, `success`, difference)

	htmlText += `</div><div class="col-1"></div></div></div></body><style>a {text-decoration: none;}* {text-align: center;}</style></html>`

	return htmlText
}

func tableText(title, titleClass string, diffs []Diff) string {
	var text = fmt.Sprintf(`<table class="table table-striped table-bordered table-hover"><thead><tr class="table-%s"><th width="10%%">行數</th><th colspan="3">%s</th></tr></thead><tbody>`, titleClass, title)
	for _, diff := range diffs {
		content1 := htmlToText(diff.Content1)
		content2 := htmlToText(diff.Content2)

		var style string
		// 只有all才用紅色顯示差別
		if titleClass == "primary" {
			if content1 != content2 {
				style = ` style="background-color: #FFE9E9;"`
			}
		}

		text += fmt.Sprintf(`<tr><td width="5%%"%s>%v</td><td>%s</td><td>%s</td></tr>`, style, diff.Idx, content1, content2)
	}
	return text + `</table>`
}

// 為了不顯示 html 語法
func htmlToText(text string) string {
	text = strings.Replace(text, `<`, `&lt;`, -1)
	return strings.Replace(text, `>`, `&gt;`, -1)
}
