package blogposts_test

import (
	"testing"

	"testing/fstest"

	blogposts "github.com/AleksaMalezic/blogposts"
)

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte("hi")},
		"hello-world2.md": {Data: []byte("hola")},
	}

	postsa := blogposts.NewPostsFromFS(fs)

	if len(postsa) != len(fs) {
		t.Errorf("got %d posts, wanted %d posts", len(postsa), len(fs))
	}
}
