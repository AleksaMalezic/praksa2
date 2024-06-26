package templating_test

import (
	"bytes"
	"testing"

	"github.com/AleksaMalezic/templating"
)

func TestRender(t *testing.T) {
	var (
		aPost = templating.Post{
			Title:       "hello world",
			Body:        "this is a post",
			Description: "this is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := templating.Render(&buf, aPost)

		if err != nil {
			t.Fatal(err)
		}

		got := buf.String()
		want := `<h1>hello world</h1>`
		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}

	})
}
