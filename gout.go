package main

import (
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/extensions"
	"log"
)

func main() {
	c := colly.NewCollector()
	extensions.RandomUserAgent(c)
	extensions.Referer(c)
	// Find and visit all links
	c.OnHTML("*", func(e *colly.HTMLElement) {
		log.Println(e)
	})

	c.OnResponse(func(response *colly.Response) {
		log.Println(response)
	})

	c.OnRequest(func(request *colly.Request) {
		log.Println(request.Headers)
	})

	c.OnError(func(response *colly.Response, err error) {
		log.Println("error", string(response.Body))
	})

	err := c.Visit("https://etherscan.io/tx/0xf5207d64e1f9948ec10a0d2ca532daf21dc9410b7b80010b8ef87af81c8796f7")
	if err != nil {
		log.Fatal(err)
	}

}
