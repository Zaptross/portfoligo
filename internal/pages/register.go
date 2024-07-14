package pages

import (
	"github.com/zaptross/portfoligo/internal/pages/blogs"
	"github.com/zaptross/portfoligo/internal/pages/projects"
)

// Blogs
var _ = registerPage(blogs.ExperienceForIndustry)

// Projects
var _ = registerPage(projects.NoRhythm)
var _ = registerPage(projects.SlayYourDragons)