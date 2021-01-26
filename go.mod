module main

go 1.15

replace siteGenerator => ./generator

require (
	github.com/go-git/go-billy/v5 v5.0.0
	github.com/go-git/go-git/v5 v5.2.0
	github.com/joho/godotenv v1.3.0 // indirect
	github.com/russross/blackfriday v1.6.0 // indirect
	siteGenerator v0.0.0-00010101000000-000000000000 // indirect
)
