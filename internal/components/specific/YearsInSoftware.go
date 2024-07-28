package specific

import (
	g "github.com/zaptross/gorgeous"
)

func YearsInSoftware() (string, g.JavaScript) {
	return "%s", g.JavaScript(`
		const duration = (Date.now() - new Date('2020-07-29').getTime()) / (365 * 24 * 60 * 60 * 1000);
		let fillText = "";

		// Fill in years
		const years = Math.floor(duration);
		fillText += years > 0 ? '%s years'.replace('%s', years) : "";

		// Fill in months
		const months = Math.floor((duration % 1) * 12);
		fillText += months > 0 ? ' and %s months'.replace('%s', months) : "";

		thisElement.textContent = thisElement.textContent.replace('%s', fillText);
	`)
}
