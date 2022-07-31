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
	var html = strings.Builder{}
	for _, line := range strings.Split(markdown, "\n") {
		if line[0] == listItemMarker {
			itemList = append(itemList, fmt.Sprintf("<li>%s</li>", renderHTML(line[2:])))
			continue
		} else if len(itemList) != 0 {
			html.WriteString(fmt.Sprintf("<ul>%s</ul>", strings.Join(itemList, "")))
			itemList = []string{}
		}
		if line[0] == headingMarker {
			headerWeight := getHeadingWeight(line)
			if headerWeight != -1 {
				html.WriteString(fmt.Sprintf("<h%d>%s</h%d>", headerWeight, line[headerWeight+1:], headerWeight))
			} else {
				html.WriteString(fmt.Sprintf("<p>%s</p>", line))
			}
			continue
		}
		html.WriteString(fmt.Sprintf("<p>%s</p>", renderHTML(line)))
	}
	if len(itemList) != 0 {
		html.WriteString(fmt.Sprintf("<ul>%s</ul>", strings.Join(itemList, "")))
	}
	return html.String()
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
