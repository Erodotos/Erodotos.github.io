package main

import (
	"fmt"
	"siteGenerator"
)

func main() {

	//Generate Site
	g := siteGenerator.New()
	posts, err := g.LoadPosts()
	if err != nil {
		fmt.Println(err)

	}

	tmpl, err := g.LoadTemplate("index.html")
	if err != nil {
		fmt.Println(err)
	}

	g.GenerateIndex(tmpl, posts)

	tmpl, err = g.LoadTemplate("post.html")
	if err != nil {
		fmt.Println(err)
	}

	g.GeneratePosts(tmpl, posts)

}
