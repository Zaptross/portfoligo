package pages

import (
	"github.com/zaptross/portfoligo/internal/types"
)

func GetPageTypeTags() []string {
	return []string{
		types.TYPE_BLOG,
		types.TYPE_PROJECT,
	}
}

func GetContentTypeTags() []string {
	return []string{
		types.TAG_GAME,
		types.TAG_SOFTWARE,
		types.TAG_DEV,
		types.TAG_DESIGN,
	}
}

func GetLangTags() []string {
	return []string{
		types.TAG_LANG_CSHARP,
		types.TAG_LANG_GO,
		types.TAG_LANG_TS,
		types.TAG_LANG_REACT,
	}
}
