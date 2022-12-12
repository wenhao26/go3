package main

import (
	"fmt"
	"strconv"

	"github.com/gocolly/colly"
)

type Document struct {
	Id      int    `json:"id"`
	Chapter string `json:"chapter"`
	Url     string `json:"url"`
}

func main() {
	c := colly.NewCollector(colly.AllowedDomains("www.readnovel.com"))

	c.OnHTML(".volume-wrap ul", func(e *colly.HTMLElement) {
		var documents []Document
		e.ForEach("li", func(i int, liE *colly.HTMLElement) {
			id, _ := strconv.Atoi(liE.Attr("data-rid"))
			chapter := liE.ChildText("a")
			url := "https://www.readnovel.com" + liE.ChildAttr("a", "href")

			document := Document{
				Id:      id,
				Chapter: chapter,
				Url:     url,
			}
			documents = append(documents, document)
			fmt.Println(id, chapter, url)
		})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Printf("Response %s: %d bytes\n", r.Request.URL, len(r.Body))
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Response %s: %d bytes\n", r.Request.URL, len(r.Body))
	})

	c.Visit("https://www.readnovel.com/book/24550170009719004#Catalog")
}
