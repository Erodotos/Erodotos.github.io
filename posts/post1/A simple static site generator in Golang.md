# A simple static site generator in Golang

## Introduction

The purpose of this article is to show how I built an elementary **static site generator** using *Golang*. I know that several open-source static **site generators** are available on the internet; however, I have decided to create my own mini-project to learn how they work and satisfy my curiosity. Furthermore, this article can be a reference point to those who might try to create their own simplistic **static site generator**.  

First of all, let us clarify some basic terminology. The term static web page refers to a web page delivered to the users' browser, as stored on a server (e.g., plain .html files). On the other hand, a  dynamic web page is generated on demand by a web application when users seek to access that page.

In the next section, I will explain the principal concepts of a **static site generator** and its major tasks. Subsequently, we will see implementation details and dive into some *Golang* code. 

## Static Site Generator : How it works

The structure of a static site generator is straightforward. Firstly, we need to prepare a template for our website and provide some data to fill in. We forward these sources into the main engine, whereas a program produces the final website according to the given parameters. We can directly host this website on a server and put our content online.

![Static Site Generator : How it works](https://cloud.netlifyusercontent.com/assets/344dbf88-fdf9-42bb-adb4-46f01eedd629/da1ef4c9-9d18-49c4-9d01-2defed1af3df/ssg-ssr-01-ssg.png)


## Implementation

For the implementation of this project, we will be using 5 technologies as shown below:

1. [Golang](#golang-site-generator)
2. [HTML](#html--css-site-templates)
3. [CSS](#html--css-site-templates)
4. [JSON](#json--markdown-site-contentssources)
5. [MArkdown](#json--markdown-site-contentssources)

### Golang (site generator)
We will be using Golang (a statically typed, compiled programming language) to implement the generator.

### HTML / CSS (site templates)
We will be using HTML and CSS to create a simple template for our website. For rapid production purposes, I recommend using Boostrap since it provides a wide variety of styles and animations.

### JSON / Markdown (site contents/sources)
We will be using JSON files to encapsulate the metadata of each article. I will be using Markdown to write the article because it is a comfortable solution accelerating the production process. The generator can use these sources (JSON file and the Markdown file ) to create the final website.

#### Step 1 - Creating a basic HTML template

HTML templates are necessary sources that will be given as input to the **static site generator**. My personal website has 2 main templates. The first regards the index.html, which presents a list of all available posts, and the second concerns the structure of a page presenting an article. The templates are designed according to the *Golang* package [template/html](https://golang.org/pkg/html/template/). To use this package, we have to include a special syntax throughout the HTML. Below you can view the syntax injected into my HTML code for both index.html and post.html templates.

&nbsp; 

The posts section below illustrates a list of articles. The list will be generated using a loop statement as shown in the brackets {{range .}} ... {{end}}.

&nbsp; 

```html
<section class="content">
    {{range .}}
        <div class="container mt-5">
            <div class="row">
                <div class="col-12">
                    <h2 class="post-title"><a href="{{.Link}}.html">{{ .Metadata.Title}}</a></h2>
                </div>
            </div>
            <div class="row">
                <div class="col-12">
                    {{.Metadata.Day_Number}} {{.Metadata.Month}} {{.Metadata.Year}} | {{.Metadata.Day_Text}}
                </div>
            </div>
        </div>
    {{end}}
</section>
```

&nbsp; 

The code below shows how Markdown files are included as a source into the website for consumption, using the *zero-md* module.

&nbsp; 

```html
<section class="content">
    <div class="container mt-5">
        <zero-md src="./{{.Metadata.Title}}.md"></zero-md>
    </div>
</section>
```

&nbsp; 

#### Step 2 - Implementing the generator

&nbsp; 

The generator performs 3 simple tasks.
+ Loads Posts from file.
+ Load Templates from file.
+ Generate website HTML files according to the Posts and Templates.

The code below depicts these tasks.

&nbsp; 


```go
func main() {

	// Create generator instance
	g := siteGenerator.New()

    // Load posts
	posts, err := g.LoadPosts()
	if err != nil {
		fmt.Println(err)
	}

    //Load index.html template
	tmpl, err := g.LoadTemplate("index.html")
	if err != nil {
		fmt.Println(err)
	}

    // Generate index.html file
	err = g.GenerateIndex(tmpl, posts)
    if err != nil {
		fmt.Println(err)
	}

    //Load post.html template
	tmpl, err = g.LoadTemplate("post.html")
	if err != nil {
		fmt.Println(err)
	}

    // Generate post.html file
	err = g.GeneratePosts(tmpl, posts)
    if err != nil {
		fmt.Println(err)
	}

}
```

&nbsp; 

The code below describes in further detail how generator functions are implemented. Also, I have included comments that allow you to understand it effortlessly.

&nbsp; 

```go
func (g *Generator) LoadPosts() ([]Post, error) {

    // Set the path for the root directory, where posts are located
	path := filepath.Join("posts")
	var postFolders []string

    // Open the root directory
	dir, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error accessing directory %s: %v", path, err)
	}
	defer dir.Close()

    // Read all post directories. Each post is encapsulated in a different folder.
	files, err := dir.Readdir(-1)
	if err != nil {
		return nil, fmt.Errorf("error reading contents of directory %s: %v", path, err)
	}

    // Append each directory to an array
	for _, file := range files {
		if file.IsDir() && file.Name()[0] != '.' {
			postFolders = append(postFolders, filepath.Join(path, file.Name()))
		}
	}

    // Read all posts from the directories and parse them into an array of Post structs
	var posts []Post
	for _, folder := range postFolders {
        // Create a new Post struct
		post, err := newPost(folder)
		if err != nil {
			return nil, fmt.Errorf("error reading post contents %s: %v", folder, err)
		}
		posts = append(posts, post)
	}

    // Sort posts according to their ID
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Metadata.Id > posts[j].Metadata.Id
	})

	return posts, nil
}
```

&nbsp; 

```go
func (g *Generator) LoadTemplate(templateName string) (*template.Template, error) {
    
    // Set the path for the directory, where the given template is located
    path := filepath.Join("templates", templateName)

    // Load the template from file and return it
	t, err := template.ParseFiles(path)
	if err != nil {
		log.Panicln(err)
		return nil, err
	}

	return t, nil
}
```

&nbsp; 

```go
func (g *Generator) GenerateIndex(t *template.Template, p []Post) error {

    // Set the export path and file name for the index.html generated code
	filePath := filepath.Join("index.html")
	f, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating file %s: %v", filePath, err)
	}

    // Create a buffer to store the generated site temporarily
    w := bufio.NewWriter(f)
    // Use the Execute() function provided by the template/html package to generate the webpage
	if err := t.Execute(w, p); err != nil {
		return fmt.Errorf("error executing template %s : %v", filePath, err)
	}

    // Flush the generated file on the disk
	if err := w.Flush(); err != nil {
		return fmt.Errorf("error writing file %s: %v", filePath, err)
	}

	f.Close()
	return nil
}
```

&nbsp; 


```go
func (g *Generator) GeneratePosts(t *template.Template, p []Post) error {

    // Receive an array of Posts and for each post generate a .html file
	for _, post := range p {
        // Set the export path and file name for the index.html generated code
		filePath := filepath.Join(post.Link + ".html")
		f, err := os.Create(filePath)
		if err != nil {
			return fmt.Errorf("Error creating file %s: %v", filePath, err)
		}

        // Create a buffer to store the generated site temporarily
        w := bufio.NewWriter(f)
        // Use the Execute() function provided by the template/html package to generate the webpage
		if err := t.Execute(w, post); err != nil {
			return fmt.Errorf("Error executing template %s : %v", filePath, err)
		}

        // Flush the generated file on the disk
		if err := w.Flush(); err != nil {
			return fmt.Errorf("Error writing file %s: %v", filePath, err)
		}
		f.Close()
	}

	return nil
}
```

#### Step 3 - Use Markdown and JSON files to write an article

&nbsp; 

As I have mentioned before, an article's body is actually represented using Markdown (a simple markup language). The picture below depicts an example of writing a document using Markdown. 

![Markdown Example](https://miro.medium.com/max/1400/0*lzRmzAy5OICef7rK.png)

Additionally, I use a simple JSON file to declare some metadata for each post. You can view below an example of metadata used to generate this article.

```json
{
    "id":1,
    "title" : "A simple static site generator in Golang",
    "day_number" : "26",
    "month" : "January",
    "year" : "2021" ,
    "day_text": "Tuesday"
}
```

&nbsp; 

## Complete Source Code of the Project [HERE](https://github.com/Erodotos/Erodotos.github.io)

&nbsp; 

## References

&nbsp; 

+ [https://en.wikipedia.org/wiki/Static_web_page](https://en.wikipedia.org/wiki/Static_web_page)
+ [https://golang.org/pkg/html/template/](https://golang.org/pkg/html/template/)
+ [https://www.markdownguide.org/basic-syntax/#code](https://www.markdownguide.org/basic-syntax/#code)
+ [https://www.zupzup.org/static-blog-generator-go/index.html](https://www.zupzup.org/static-blog-generator-go/index.html)