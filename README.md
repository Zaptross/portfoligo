# Portfoligo

Portfoligo is a personal blog and portfolio website, written in Go using the `zaptross/gorgeous` library for multi-page static site generation.

## Building

To run Portfoligo locally, clone this repository and ensure Go is installed, before running the following commands from the root of the repo:

```bash
# install dependencies
go mod download

# run the code generation
cd ./gen/
go run ./

# build the static site
cd ./cmd/
go run ./
```

## Watched Builds

To watch for changes and rebuild the site automatically, run the following commands:

```bash
./watch.sh
```

This will watch for changes in the `./internal` directory and rebuild the site automatically using `inotifywait`.

To automatically reload the browser on changes, the 'Live Server' extension for Visual Studio Code is recommended.

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

Notably, there exists a utility type `pages.TYPE_ROOT = "root"` which will result in a page being generated as the index of it's slug type, eg:

```go
page := types.PageDetails{
    Type: pages.TYPE_ROOT,
    Slug: "blog",
    // ... etc.
}
```

Will result in a page being generated at `dist/blog/index.html`.

For organisation, the `./internal/pages` directory is split into subdirectories for each page type, eg: `blog`, `projects`

## Patterns

The `internal/components` directory contains reusable components in the style of [Atomic Design](https://bradfrost.com/blog/post/atomic-web-design/), which can be used to build up pages in `internal/pages`.

The `internal/theme` directory contains [Theme UI](https://theme-ui.com/) inspired tokenised theme definitions, which can be used to style the site.
