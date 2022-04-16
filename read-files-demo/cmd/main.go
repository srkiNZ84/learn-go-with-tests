package main

import (
	"fmt"
	"log"
	"os"

	blogposts "serge.sh"
)

func main() {
	posts, err := blogposts.PostsFromFS(os.DirFS("posts"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(posts)
}
