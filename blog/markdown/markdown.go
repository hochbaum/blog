package markdown

import (
	"github.com/russross/blackfriday"
	"html/template"
)

// ToHtml converts markdown-formatted text to HTML using the blackfriday library.
func ToHtml(text string) template.HTML {
	return template.HTML(blackfriday.Run([]byte(text)))
}
