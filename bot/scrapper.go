package bot

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

func ScraperSet50(url string) ([]SSetTrade, error) {
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
						// fmt.Println("Error: Expected 5 lines of data")
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

	// fmt.Println("data...", data)
	fmt.Println("stopping bot...")
	return data, nil
}

func ScraperGold(url string) ([]SGold, error) {
	fmt.Println("running bot...")
	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	var data []SGold
	c.OnHTML("tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(trIdx int, tr *colly.HTMLElement) {
			if trIdx < 2 { // Process only the first three rows
				var raw SGold
				tr.ForEach("td", func(tdIdx int, td *colly.HTMLElement) {
					switch tdIdx {
					case 0:
						raw.Type = td.Text
					case 1:
						raw.Buy = td.Text
					case 2:
						raw.Sell = td.Text
					}
				})
				if raw.Type == "ทองคำแท่ง" {
					data = append(data, raw)
				} else if raw.Type == "ทองรูปพรรณ" {
					data = append(data, raw)
				}
			}
		})
	})
	err := c.Visit(url)
	if err != nil {
		fmt.Println("Error scraping:", err)
	}

	// fmt.Println("data...", data)
	fmt.Println("stopping bot...")
	return data, nil
}

func ScraperGold2(url string) ([]SGold, error) {
	fmt.Println("running bot...")
	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	var data []SGold
	c.OnHTML("tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(trIdx int, tr *colly.HTMLElement) {
			if trIdx < 2 { // Process only the first three rows
				var raw SGold
				tr.ForEach("td", func(tdIdx int, td *colly.HTMLElement) {
					switch tdIdx {
					case 0:
						raw.Type = td.Text
					case 1:
						raw.Buy = td.Text
					case 2:
						raw.Sell = td.Text
					}
				})
				data = append(data, raw)
				// if raw.Type == "ทองคำแท่ง" {
				// 	data = append(data, raw)
				// } else if raw.Type == "ทองรูปพรรณ" {
				// 	data = append(data, raw)
				// }

			}
		})
	})
	err := c.Visit(url)
	if err != nil {
		fmt.Println("Error scraping:", err)
	}

	// ================================
	rand.Seed(time.Now().UnixNano())

	// Generate a random number
	num := rand.Int()

	// Set error based on the odd/even check
	if num%2 == 0 {
		err = nil
	} else {
		err = errors.New("Test error")
	}

	// // fmt.Println("data...", data)
	// fmt.Println("stopping bot...")
	return data, err
}
