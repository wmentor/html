package html

import (
	"bytes"
	"io"
	"strings"

	ht "golang.org/x/net/html"
)

type HTML struct {
	links     map[string]bool
	iframes   map[string]bool
	eraseCnt  int
	output    *bytes.Buffer
	lastEol   bool
}

var (
	tagErase map[string]bool
)

func init() {

	tagErase = map[string]bool{}

	for _, tg := range []string{"form", "iframe", "link", "meta", "noscript", "option", "script", "select", "style"} {
		tagErase[tg] = true
	}

}

func New() *HTML {
	res := &HTML{}

	res.links = make(map[string]bool)
	res.iframes = make(map[string]bool)

	res.output = bytes.NewBuffer(nil)

	return res
}

func (h *HTML) ParseString(str string) {
	h.Parse(strings.NewReader(str))
}

func (h *HTML) beforeParse() {

	h.links = make(map[string]bool)
	h.iframes = make(map[string]bool)
	
	h.eraseCnt = 0
	h.lastEol = true

	h.output.Reset()
}

func (h *HTML) Parse(r io.Reader) {

	h.beforeParse()

	parser := ht.NewTokenizer(r)

	for {
		tt := parser.Next()

		switch {

		case tt == ht.ErrorToken:
			return

		case tt == ht.StartTagToken:
			t := parser.Token()
			h.onStartTag(&t)

		case tt == ht.EndTagToken:
			t := parser.Token()
			h.onCloseTag(&t)

		case tt == ht.SelfClosingTagToken:
			t := parser.Token()
			h.onStartTag(&t)
			h.onCloseTag(&t)

		case tt == ht.TextToken:
			h.onText(parser.Text())

		}
	}
}

func (h *HTML) onStartTag(t *ht.Token) {

	if tagErase[t.Data] {
		if t.Data != "meta" && t.Data != "link" {
			h.eraseCnt++
		}
	}

	switch t.Data {
	
	case "iframe":
		
		for _, a := range t.Attr {
			if a.Key == "src" {
				h.iframes[a.Val] = true
			}
		}
	
	case "a":

		for _, a := range t.Attr {
			if a.Key == "href" {
				h.links[a.Val] = true
			}
		}
	}

	if h.eraseCnt > 0 {
		return
	}

	if t.Data == "br" {
		if !h.lastEol {
			h.output.WriteRune('\n')
			h.lastEol = true
		}
	}
}

func (h *HTML) onCloseTag(t *ht.Token) {

	if tagErase[t.Data] && t.Data != "meta" && t.Data != "link" {
		h.eraseCnt--
		if h.eraseCnt < 0 {
			h.eraseCnt = 0
		}
	}

	if h.eraseCnt > 0 {
		return
	}

	if t.Data == "p" || t.Data == "div" || t.Data == "ol" || t.Data == "ul" {
		if !h.lastEol {
			h.output.WriteRune('\n')
			h.lastEol = true
		}
	}
}

func (h *HTML) onText(data []byte) {
	if h.eraseCnt == 0 {

		data = bytes.TrimSpace(data)
		if len(data) == 0 {
			return
		}

		if !h.lastEol {
			h.output.WriteRune(' ')
		}
		h.lastEol = false
		h.output.Write(data)
	}
}

func (h *HTML) Text() []byte {
	return h.output.Bytes()
}