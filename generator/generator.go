package siteGenerator

import (
	"bufio"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
)

type Generator struct {
}

type Post struct {
	Metadata struct {
		Id         int    `json:"id"`
		Title      string `json:"title"`
		Day_Number string `json:"day_number"`
		Month      string `json:"month"`
		Year       string `json:"year"`
		Day_Text   string `json:"day_text"`
	}
	Link string
	// ContentsLink template.HTML
}

func New() *Generator {
	return &Generator{}
}

func newPost(path string) (Post, error) {

	post := Post{}

	filePath := filepath.Join(path, "metadata.json")
	metadata, err := ioutil.ReadFile(filePath)
	if err != nil {
		return post, fmt.Errorf("error while reading file %s: %v", filePath, err)
	}
	err = json.Unmarshal(metadata, &post.Metadata)
	if err != nil {
		return post, fmt.Errorf("error reading yml in %s: %v", filePath, err)
	}

	// filePath = filepath.Join(path, "contents.md")
	// contents, err := ioutil.ReadFile(filePath)
	// if err != nil {
	// 	return post, fmt.Errorf("error while reading file %s: %v", filePath, err)
	// }
	// html := blackfriday.MarkdownCommon(contents)
	// if err != nil {
	// 	return post, fmt.Errorf("error during syntax highlighting of %s: %v", filePath, err)
	// }
	// post.Contents = template.HTML(string(html))

	post.Link = "./posts/post" + strconv.Itoa(post.Metadata.Id) + "/" + post.Metadata.Title

	return post, nil
}

// Load Posts from file

func (g *Generator) LoadPosts() ([]Post, error) {

	path := filepath.Join("posts")
	var postFolders []string

	dir, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error accessing directory %s: %v", path, err)
	}
	defer dir.Close()

	files, err := dir.Readdir(-1)
	if err != nil {
		return nil, fmt.Errorf("error reading contents of directory %s: %v", path, err)
	}

	for _, file := range files {
		if file.IsDir() && file.Name()[0] != '.' {
			postFolders = append(postFolders, filepath.Join(path, file.Name()))
		}
	}

	var posts []Post
	for _, folder := range postFolders {
		post, err := newPost(folder)
		if err != nil {
			return nil, fmt.Errorf("error reading post contents %s: %v", folder, err)
		}
		posts = append(posts, post)
	}

	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Metadata.Id < posts[j].Metadata.Id
	})

	return posts, nil
}

// Load template

func (g *Generator) LoadTemplate(templateName string) (*template.Template, error) {
	path := filepath.Join("templates", templateName)

	t, err := template.ParseFiles(path)
	if err != nil {
		log.Panicln(err)
		return nil, err
	}

	return t, nil
}

// Fill template with posts

func (g *Generator) GenerateIndex(t *template.Template, p []Post) error {
	filePath := filepath.Join("index.html")
	f, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating file %s: %v", filePath, err)
	}

	w := bufio.NewWriter(f)
	if err := t.Execute(w, p); err != nil {
		return fmt.Errorf("error executing template %s : %v", filePath, err)
	}

	if err := w.Flush(); err != nil {
		return fmt.Errorf("error writing file %s: %v", filePath, err)
	}

	f.Close()
	return nil
}

// Create individual post pages

func (g *Generator) GeneratePosts(t *template.Template, p []Post) error {

	for _, post := range p {
		fmt.Println(post.Link)
		filePath := filepath.Join(post.Link + ".html")
		f, err := os.Create(filePath)
		if err != nil {
			fmt.Println(1, err)
			return fmt.Errorf("Error creating file %s: %v", filePath, err)
		}

		w := bufio.NewWriter(f)
		if err := t.Execute(w, post); err != nil {
			fmt.Println(2, err)
			return fmt.Errorf("Error executing template %s : %v", filePath, err)
		}

		if err := w.Flush(); err != nil {
			return fmt.Errorf("Error writing file %s: %v", filePath, err)
		}
		f.Close()
	}

	return nil
}
