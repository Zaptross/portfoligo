package pages

import (
	industryproject "github.com/zaptross/portfoligo/internal/pages/blogs/industry-project"
	"github.com/zaptross/portfoligo/internal/pages/projects"
)

// Blogs
var _ = registerPage(industryproject.ExperienceForIndustry)
var _ = registerPage(industryproject.Beginnings)
var _ = registerPage(industryproject.Redesign)
var _ = registerPage(industryproject.Server)
var _ = registerPage(industryproject.Backburner)
var _ = registerPage(industryproject.Recap)

// Projects
var _ = registerPage(projects.NoRhythm)
var _ = registerPage(projects.SlayYourDragons)
var _ = registerPage(projects.PlayerLikeMechAi)
var _ = registerPage(projects.AshesOfTheVeil)
var _ = registerPage(projects.AtAnyCost)
var _ = registerPage(projects.KowloonExigency)
