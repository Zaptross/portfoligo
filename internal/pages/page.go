package pages

import (
	"fmt"

	"github.com/samber/lo"
	"github.com/zaptross/portfoligo/internal/types"
)

var allPages []types.PageDetails

func registerPage(newPage types.PageDetails) types.PageDetails {
	_, exists := lo.Find(
		allPages,
		func(p types.PageDetails) bool { return p.Slug == newPage.Slug },
	)
	if exists {
		panic(fmt.Sprintf("Page with slug %s already exists", newPage.Slug))
	}

	allPages = append(allPages, newPage)

	return newPage
}

func GetAllPagesByType(t string) []types.PageDetails {
	return lo.Filter(
		allPages,
		func(p types.PageDetails, _ int) bool {
			return p.Type == t
		},
	)
}

func GetAllPagesByTypes() map[string][]types.PageDetails {
	return map[string][]types.PageDetails{
		TYPE_ROOT:    GetAllPagesByType(TYPE_ROOT),
		TYPE_BLOG:    GetAllPagesByType(TYPE_BLOG),
		TYPE_PROJECT: GetAllPagesByType(TYPE_PROJECT),
	}
}
