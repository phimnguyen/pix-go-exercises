package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

const urlPrefix = "https://vk.com"

func scrape() {
	res, err := http.Get("https://vk.com/doujinmusic")
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".wall_posts").Children().Each(func(i int, s *goquery.Selection) {
		postID, _ := s.Find(".wi_date").Attr("href")

		post := s.Find(`a[class*="medias_link"]`)
		album := post.Text()

		link, exists := post.Attr("href")
		if exists == false {
			return
		}
		fmt.Println(urlPrefix + postID)
		fmt.Println(fmt.Sprintf("%s\n%s%s\n", album, urlPrefix, link))
	})
}

func main() {
	scrape()
}
