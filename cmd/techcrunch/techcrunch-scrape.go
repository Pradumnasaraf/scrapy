package techcrunch

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/alexeyco/simpletable"
	"github.com/Pradumnasaraf/scrapy/helpers"
)

var cells [][]*simpletable.Cell

func techcrunch(category string) {

	fmt.Println("Searching for latest blogs on techcrunch....")

	res, err := getResponse(category)
	if err != nil {
		log.Fatal(err)
	}

	techCrunchScrape(res)

	createTable()

}

func techCrunchScrape(res *http.Response) {

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("div.river").Find("div.post-block").Each(func(i int, item *goquery.Selection) {
		title := strings.TrimSpace(item.Find("h2.post-block__title").Text())
		author := strings.TrimSpace(item.Find("span.river-byline__authors").Find("a").Text())
		time := strings.TrimSpace(item.Find("time.river-byline__time").Text())
		url := strings.TrimSpace(item.Find("h2.post-block__title").Find("a").AttrOr("href", ""))

		if len(author) > 15 {
			author = author[:14] + ".."
		}

		index := helpers.Purple(strconv.Itoa(i + 1))

		url = helpers.Green(url)

		appendTableData(title, author, time, url, index)

	})

}

func appendTableData(title, author, time, urls, index string) {

	cell := []*simpletable.Cell{
		{Align: simpletable.AlignLeft, Text: index},
		{Align: simpletable.AlignLeft, Text: title + "\n URL: " + urls},
		{Align: simpletable.AlignLeft, Text: author},
		{Align: simpletable.AlignLeft, Text: time},
	}
	cells = append(cells, cell)

}

func createTable() {

	table := simpletable.New()

	// Set the headers
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Title"},
			{Align: simpletable.AlignCenter, Text: "Author"},
			{Align: simpletable.AlignCenter, Text: "Time"},
		},
	}

	// Table Body
	table.Body = &simpletable.Body{Cells: cells}

	// Set the style
	table.SetStyle(simpletable.StyleUnicode)
	fmt.Println(table.String())

}