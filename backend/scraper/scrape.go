package scraper

import (
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

const numberOfItems = 5

type ScrapedItem struct {
	Name     string `json:"name"`
	Price    string `json:"price"`
	ImageUrl string `json:"imageUrl"`
	Url      string `json:"url"`
}

func Scrape(title string, category string, brand string) (*[]ScrapedItem, error) {
	c := colly.NewCollector()
	c.SetRequestTimeout(120 * time.Second)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})

	results := make([]ScrapedItem, 0)
	c.OnHTML("div.product-info", func(h *colly.HTMLElement) {
		var item ScrapedItem
		item.Name = h.DOM.Find("h1.page-title").Text()
		h.DOM.Find("span.price-container").Children().Each(func(i int, s *goquery.Selection) {
			priceType, ptexists := s.Attr("data-price-type")
			if ptexists && priceType == "finalPrice" {
				item.Price, _ = s.Attr("data-price-amount")
			}
		})
		item.Url = h.Request.URL.String()
		//TODO: fetch image as well
		results = append(results, item)
	})
	c.OnHTML("ol.products", func(e *colly.HTMLElement) {
		e.ForEachWithBreak("li", func(i int, h *colly.HTMLElement) bool {
			if i < numberOfItems {
				var item ScrapedItem
				li := h.DOM
				name := li.Find("strong.product").Text()
				imageUrl, imageExists := li.Find("img.product-image-photo").Attr("src")
				if imageExists {
					item.ImageUrl = imageUrl
				} else {
					item.ImageUrl = ""
				}
				var price string
				li.Find("span.price-container").Children().Each(func(i int, s *goquery.Selection) {
					priceType, ptexists := s.Attr("data-price-type")
					if ptexists && priceType == "finalPrice" {
						price, _ = s.Attr("data-price-amount")
					}
				})
				if url, urlExists := li.Find("a.product.photo.product-item-photo").Attr("href"); urlExists {
					item.Url = url
				}
				item.Price = price
				item.Name = name
				results = append(results, item)
				return true
			} else {
				return false
			}
		})
	})

	c.Visit(fmt.Sprint("https://www.sastodeal.com/catalogsearch/result/?q=", url.QueryEscape(title)))
	if len(results) == 0 && brand != "" {
		c.Visit(fmt.Sprint("https://www.sastodeal.com/catalogsearch/result/?q=", url.QueryEscape(brand)))
	}
	if len(results) == 0 {
		c.Visit(fmt.Sprint("https://www.sastodeal.com/catalogsearch/result/?q=", url.QueryEscape(category)))
	}
	return &results, nil
}

func TrimTitle(title string) string {
	if length := len(strings.Fields(title)); length <= 5 {
		return title
	} else if length > 8 {
		shortened := strings.Join(strings.Split(title, " ")[:5], " ")
		return shortened
	}
	return title
}
