package markdown

// implementation to refactor

import (
	"fmt"
	"strings"
)

func Render(markdown string) string {
	var itemList []string
	var html string
	for _, line := range strings.SplitN(markdown, "\n", -1) {
		if line[0] == '*' {
			itemList = append(itemList, "<li>"+renderHTML(line[2:])+"</li>")
			continue
		} else if len(itemList) != 0 {
			html += printList(&itemList)
		}
		if line[0] == '#' {
			headerWeight := checkWeight(line)
			if headerWeight >= 1 && headerWeight <= 6 {
				html += fmt.Sprintf("<h%d>%s</h%d>", headerWeight, line[headerWeight+1:], headerWeight)
			} else {
				html += fmt.Sprintf("<p>%s</p>", line)
			}
			continue
		}
		html += "<p>" + renderHTML(line) + "</p>"
	}
	html += printList(&itemList)
	return html
}

var headingWeights = []string{"######", "#####", "####", "###", "##", "#"}

func checkWeight(line string) int {
	if strings.HasPrefix(line, "#######") {
		return -1
	}
	for _, prefix := range headingWeights {
		if strings.HasPrefix(line, prefix) {
			return len(prefix)
		}
	}
	return -1
}

func renderHTML(markdownLine string) string {
	htmlLine := markdownLine
	for strings.Contains(htmlLine, "__") {
		htmlLine = strings.Replace(htmlLine, "__", "<strong>", 1)
		htmlLine = strings.Replace(htmlLine, "__", "</strong>", 1)
	}
	for strings.Contains(htmlLine, "_") {
		htmlLine = strings.Replace(htmlLine, "_", "<em>", 1)
		htmlLine = strings.Replace(htmlLine, "_", "</em>", 1)
	}
	return htmlLine
}

func printList(itemList *[]string) string {
	defer func() {
		*itemList = []string{}
	}()
	if len(*itemList) != 0 {
		return "<ul>" + strings.Join(*itemList, "") + "</ul>"
	}
	return ""
}
