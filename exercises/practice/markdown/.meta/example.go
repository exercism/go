package markdown

import (
	"fmt"
	"strings"
)

const (
	headingMarker  = '#'
	listItemMarker = '*'
)

// Render translates markdown to HTML
func Render(markdown string) string {
	var itemList []string
	var html string
	for _, line := range strings.Split(markdown, "\n") {
		if line[0] == listItemMarker {
			itemList = append(itemList, fmt.Sprintf("<li>%s</li>", renderHTML(line[2:])))
			continue
		} else if len(itemList) != 0 {
			html += printList(&itemList)
		}
		if line[0] == headingMarker {
			headerWeight := getHeadingWeight(line)
			if headerWeight != -1 {
				html += fmt.Sprintf("<h%d>%s</h%d>", headerWeight, line[headerWeight+1:], headerWeight)
			} else {
				html += fmt.Sprintf("<p>%s</p>", line)
			}
			if headerWeight <= 6 {

			} else {

			}
			continue
		}
		html += "<p>" + renderHTML(line) + "</p>"
	}
	html += printList(&itemList)
	return html
}

func getHeadingWeight(line string) int {
	for i := 0; i <= 6; i++ {
		if line[i] != headingMarker {
			return i
		}
	}
	return -1
}

func renderHTML(markdownLine string) string {
	htmlLine := markdownLine
	// Convert all __ pairs into <strong>...</strong>
	for strings.Contains(htmlLine, "__") {
		htmlLine = strings.Replace(htmlLine, "__", "<strong>", 1)
		htmlLine = strings.Replace(htmlLine, "__", "</strong>", 1)
	}
	// Convert all _ pairs into <em>...</em>
	for strings.Contains(htmlLine, "_") {
		htmlLine = strings.Replace(htmlLine, "_", "<em>", 1)
		htmlLine = strings.Replace(htmlLine, "_", "</em>", 1)
	}
	return htmlLine
}

func printList(itemList *[]string) string {
	// empty list after return
	defer func() {
		*itemList = []string{}
	}()
	if len(*itemList) != 0 {
		return fmt.Sprintf("<ul>%s</ul>", strings.Join(*itemList, ""))
	}
	// reset list
	return ""
}
