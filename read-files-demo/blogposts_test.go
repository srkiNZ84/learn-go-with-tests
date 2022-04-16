package blogposts_test

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"

	blogposts "serge.sh"
)

func TestBlogPost(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {

		fileSystem := fstest.MapFS{
			"hello-world.md": {Data: []byte(`Title: Hello, TDD world!
Description: lol
Tags: tdd, go
---
Hello
World`)},
			"hello-world2.md": {Data: []byte(`Title: Hello, TDD world2!
Description: lol2
Tags: tdd2, go2
---
Hello
World2`)},
			//"hello-twitch.md": {Data: []byte("Title: Hello twitchy world!")}, ???
		}
		posts, err := blogposts.PostsFromFS(fileSystem)

		if err != nil {
			t.Fatal(err)
		}

		if len(posts) != len(fileSystem) {
			t.Errorf("Expected %d posts, got %d posts", len(fileSystem), len(posts))
		}

		assertPost(t, posts[0], blogposts.Post{
			Title:       "Hello, TDD world!",
			Description: "lol",
			Tags:        []string{"tdd", "go"},
			Body: `Hello
World`,
		})
	})

	t.Run("failing file system", func(t *testing.T) {
		_, err := blogposts.PostsFromFS(FailingFS{})
		if err == nil {
			t.Error("expected an error but got none")
		}
	})
}

func assertPost(t *testing.T, got, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %#v, want %#v", got, want)
	}
}

type FailingFS struct {
}

func (f FailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no, I always fail")
}
