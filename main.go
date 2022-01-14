package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
)

func main() {
	c := colly.NewCollector(
		colly.Async(true),
		colly.UserAgent("Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"),
	)

	c.Limit(&colly.LimitRule{DomainGlob: "*.douban.*", Parallelism: 5})

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	//c.OnHTML(".hd", func(e *colly.HTMLElement) {
	//	log.Println(strings.Split(e.ChildAttr("a", "href"), "/")[4],
	//		strings.TrimSpace(e.DOM.Find("span.title").Eq(0).Text()))
	//})

	c.OnHTML(".tab-list", func(e *colly.HTMLElement) {

		fmt.Println(e.DOM.Find("span").Eq(0).Text())
		fmt.Println(e.DOM.Find("a").Eq(0).Text())
		fmt.Println(e.DOM.Find("a").Attr("href"))
		//fmt.Println(e.DOM.Find("span").Eq(0).Text())
	})

	c.OnHTML(".paginator a", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.Visit("https://china.nba.cn/playerindex/")
	c.Wait()
}
