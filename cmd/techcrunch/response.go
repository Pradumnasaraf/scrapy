package techcrunch

import (
	"fmt"
	"net/http"
)

func getResponse(category string) (*http.Response, error) {

	defaultUrl := "https://techcrunch.com/"
	categoryUrl := ""

	if category == "" {
		categoryUrl = defaultUrl
	} else {
		categoryUrl = getCategoryUrl(category)

		if categoryUrl == defaultUrl && category != "" {
			fmt.Println("Invalid category. Showing homepage instead.")
		}
	}

	res, err := http.Get(categoryUrl)
	if err != nil {
		return nil, err
	}
	return res, nil

}

func getCategoryUrl(category string) string {

	var url string

	switch category {
	case "startups":
		url = "https://techcrunch.com/category/startups/"
	case "apps":
		url = "https://techcrunch.com/category/apps/"
	case "ai", "artificial intelligence":
		url = "https://techcrunch.com/category/artificial-intelligence/"
	case "crypto", "cryptocurrency":
		url = "https://techcrunch.com/category/cryptocurrency/"
	case "security":
		url = "https://techcrunch.com/category/security/"
	case "gadgets":
		url = "https://techcrunch.com/category/gadgets/"
	case "venture", "venture capital", "vc", "funding":
		url = "https://techcrunch.com/category/venture-capital/"
	case "social":
		url = "https://techcrunch.com/category/social/"
	default:
		return "https://techcrunch.com/"
	}
	return url

}
