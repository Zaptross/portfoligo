package pages

const (
	// Type tags
	TYPE_BLOG    = "blog"
	TYPE_PROJECT = "project"
	TYPE_ROOT    = "root"

	// Content type tags
	TAG_GAME     = "game"
	TAG_SOFTWARE = "software"
	TAG_WEB      = "web"
	TAG_DEV      = "development"
	TAG_DESIGN   = "design"

	// Langs
	TAG_LANG_CSHARP = "csharp"
	TAG_LANG_GO     = "go"
	TAG_LANG_TS     = "typescript"
	TAG_LANG_REACT  = "react"
	TAG_LANG_RUST   = "rust"

	// Meta tags
	TAG_AVAILABLE = "available"
)

func GetPageTypeTags() []string {
	return []string{
		TYPE_BLOG,
		TYPE_PROJECT,
	}
}

func GetContentTypeTags() []string {
	return []string{
		TAG_GAME,
		TAG_SOFTWARE,
		TAG_DEV,
		TAG_DESIGN,
	}
}

func GetLangTags() []string {
	return []string{
		TAG_LANG_CSHARP,
		TAG_LANG_GO,
		TAG_LANG_TS,
		TAG_LANG_REACT,
	}
}
