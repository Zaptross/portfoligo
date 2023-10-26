# Portfoligo

Portfoligo is a personal blog and portfolio website, written in Go using the `zaptross/gorgeous` library for multi-page static site generation.

## Building

To run Portfoligo locally, clone this repository and ensure Go is installed, before running the following commands from the root of the repo:

```bash
# install dependencies
go mod download

# build the static site
cd ./cmd/
go run ./
```

## File Structure

New pages can be added by creating a new `.go` file in the `./internal/pages` directory.
The `Type` and `Slug` properties of the page will inform the file structure and url of the resulting page, eg:

```go
page := types.PageDetails{
    Type: "blog",
    Slug: "my-first-post",
    // ... etc.
}
```

Will result in a page being generated at `dist/blog/my-first-post/index.html`.
