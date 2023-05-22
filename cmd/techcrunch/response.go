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
	case "startups", "startup", "entrepreneur", "entrepreneurship":
		url = "https://techcrunch.com/category/startups/"
	case "apps", "app", "application", "applications":
		url = "https://techcrunch.com/category/apps/"
	case "ai", "artificial intelligence":
		url = "https://techcrunch.com/category/artificial-intelligence/"
	case "crypto", "cryptocurrency":
		url = "https://techcrunch.com/category/cryptocurrency/"
	case "security", "cybersecurity", "cyber security":
		url = "https://techcrunch.com/category/security/"
	case "gadgets", "gadget":
		url = "https://techcrunch.com/category/gadgets/"
	case "venture", "venture capital", "vc", "funding":
		url = "https://techcrunch.com/category/venture/"
	case "social", "social media", "social networking":
		url = "https://techcrunch.com/category/social/"
	case "fintech", "finance", "financial":
		url = "https://techcrunch.com/category/fintech/"
	case "transportation", "transport", "self-driving", "self driving":
		url = "https://techcrunch.com/category/transportation/"
	case "space", "space tech", "space technology":
		url = "https://techcrunch.com/category/space/"
	case "enterprise", "enterprise software":
		url = "https://techcrunch.com/category/enterprise/"
	case "commerce", "e-commerce", "ecommerce":
		url = "https://techcrunch.com/category/commerce/"
	case "climate", "climate tech", "climate technology":
		url = "https://techcrunch.com/category/climate/"
	case "health", "biotech", "biotechnology":
		url = "https://techcrunch.com/category/biotech-health/"
	case "games", "gaming", "game", "video games", "video game":
		url = "https://techcrunch.com/category/gaming/"
	case "hardware", "hardware startup", "hardware startups":
		url = "https://techcrunch.com/category/hardware/"
	case "robotics", "robot", "robots":
		url = "https://techcrunch.com/category/robotics/"
	case "privacy", "data privacy":
		url = "https://techcrunch.com/category/privacy/"
	case "policy", "government", "politics":
		url = "https://techcrunch.com/category/policy/"
	case "media", "entertainment":
		url = "https://techcrunch.com/category/media-entertainment/"
	default:
		return "https://techcrunch.com/"
	}
	return url

}
