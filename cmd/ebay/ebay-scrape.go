package ebay

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func ebay() {

	ebayScrape()

}

func getResponse(url string) (*http.Response, error) {

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return res, nil

}

func ebayScrape() {

	res, err := getResponse("https://www.ebay.com/sch/i.html?_from=R40&_nkw=apple&_pgn=1")
	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("li.s-item__pl-on-bottom").Each(func(i int, item *goquery.Selection) {
		title := item.Find("div.s-item__title").Find("span").Text()
		price := item.Find("span.s-item__price").Text()
		log.Println(title, price)
	})

}
