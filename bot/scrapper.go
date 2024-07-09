package bot

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

type SSetTrade struct {
	Index  string
	Price  string
	Change string
	Volume string
	Value  string
}

func Scraper(url string) ([]SSetTrade, error) {
	fmt.Println("running bot...")
	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	var data []SSetTrade
	c.OnHTML("div.table-market-overview", func(e *colly.HTMLElement) {
		e.ForEach("tbody", func(idx int, el *colly.HTMLElement) {
			if idx == 0 {
				// e.ForEach("tr", func(idx int, el *colly.HTMLElement) {
				// 	if idx == 11 {
				// 		fmt.Println("Th text:", idx, el.DOM.Children().Text())
				// 	}
				// })
				e.ForEach("tr", func(idx int, el *colly.HTMLElement) {
					d := el.Text

					lines := strings.Split(d, "\n")
					var cleanedLines []string
					for _, line := range lines {
						trimmedLine := strings.TrimSpace(line)
						if trimmedLine != "" {
							cleanedLines = append(cleanedLines, trimmedLine)
						}
					}
					if len(cleanedLines) != 5 {
						fmt.Println("Error: Expected 5 lines of data")
						return
					}
					s := SSetTrade{
						Index:  cleanedLines[0],
						Price:  cleanedLines[1],
						Change: cleanedLines[2],
						Volume: cleanedLines[3],
						Value:  cleanedLines[4],
					}
					data = append(data, s)
				})
			}

		})
	})
	err := c.Visit(url)
	if err != nil {
		fmt.Println("Error scraping:", err)
	}

	fmt.Println("data...", data)
	fmt.Println("stopping bot...")
	return data, nil
}
