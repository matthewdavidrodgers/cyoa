package adventure

import (
	"bytes"
	"encoding/json"
	"html/template"
	"io/ioutil"
)

type StoryNodePath struct {
	Text     string `json:"text"`
	NodeName string `json:"arc"`
}

type StoryNode struct {
	Title string          `json:"title"`
	Story []string        `json:"story"`
	Paths []StoryNodePath `json:"options"`
}

type Story map[string]StoryNode

var templates *template.Template

type loaderError struct{}

func (e *loaderError) Error() string {
	return "Could not load json"
}

func Load(path string) (Story, error) {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, &loaderError{}
	}
	s := make(Story)
	err = json.Unmarshal(raw, &s)
	if err != nil {
		return nil, &loaderError{}
	}
	templates, err = template.ParseFiles(
		"./adventure/page.html",
		"./adventure/node.html",
	)
	if err != nil {
		return nil, &loaderError{}
	}
	return s, nil
}

func RenderPage(node StoryNode) ([]byte, error) {
	var buf bytes.Buffer
	err := templates.ExecuteTemplate(&buf, "page", node)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
