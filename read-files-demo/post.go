package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

const (
	titlePrefix       = "Title: "
	descriptionPrefix = "Description: "
	tagsPrefix        = "Tags: "
)

func newPost(blogFile io.Reader) Post {
	scanner := bufio.NewScanner(blogFile)

	readLine := func(prefix string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), prefix)
	}
	title := readLine(titlePrefix)
	description := readLine(descriptionPrefix)
	tags := strings.Split(readLine(tagsPrefix), ", ")
	body := readBody(scanner)
	return Post{Title: title, Description: description, Tags: tags, Body: body}
}

func readBody(scanner *bufio.Scanner) string {
	scanner.Scan() // skip hyphens
	var buf bytes.Buffer
	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}
	return strings.TrimSuffix(buf.String(), "\n")
}
