package markdown

import (
	"bytes"
	"fmt"
	"regexp"
)

type charType int

// used to keep track of which tag to close
const (
	none charType = iota
	hash
	star
)

// Render translates markdown to HTML
func Render(markdown string) (html string) {
	//first the easy one, via regexp substitution
	reStrong := regexp.MustCompile("(__)(.*)(__)")
	s := reStrong.ReplaceAll([]byte(markdown), []byte("<strong>$2</strong>"))
	reEm := regexp.MustCompile("(_)(.*)(_)")
	s = reEm.ReplaceAll(s, []byte("<em>$2</em>"))
	//now manage <li> and <hN>
	var output bytes.Buffer
	starcount := 0
	hcount := 0
	needtoClose := none
	i := 0
	for i < len(s) {
		switch s[i] {
		case '#':
			for s[i] == '#' {
				hcount++
				i++
			}
			output.WriteString(fmt.Sprintf("<h%d>", hcount))
			needtoClose = hash
		case '*':
			if starcount == 0 {
				output.WriteString("<ul>")
			}
			i++
			starcount++
			output.WriteString("<li>")
			needtoClose = star
		case '\n':
			if needtoClose == hash {
				output.WriteString(fmt.Sprintf("</h%d>", hcount))
			}
			if needtoClose == star {
				output.WriteString("</li>")
			}

		default:
			output.WriteByte(s[i])

		}
		i++
	}
	if starcount > 0 || hcount > 0 {
		if needtoClose == hash {
			output.WriteString(fmt.Sprintf("</h%d>", hcount))
		}
		if needtoClose == star {
			output.WriteString("</li></ul>")
		}
		return output.String()
	}
	return fmt.Sprintf("<p>%s</p>", string(s))
}
