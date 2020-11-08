package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

//GetComicURL Test comment
func GetComicURL() string {

	return scrapeWebPage()
}

func scrapeWebPage() string {

	randDate := randomDate()

	url := fmt.Sprintf("https://www.gocomics.com/garfield/%s", randDate)

	var comicURL string

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".item-comic-image").Each(func(i int, s *goquery.Selection) {
		imageNode := s.Find(".img-fluid")
		str, exists := imageNode.Attr("src")
		if exists {
			comicURL = fmt.Sprintf("%s.png", str)
		}
	})

	return comicURL
}

func randomDate() string {

	minDate := time.Date(1978, 6, 19, 0, 0, 0, 0, time.UTC).Unix()
	maxDate := time.Now().Unix()

	delta := maxDate - minDate

	sec := rand.Int63n(delta) + minDate

	utcDate := time.Unix(sec, 0)

	randomDate := utcDate.UTC().Format("2006/01/02")

	formattedDateString := fmt.Sprintf("%s", randomDate)

	return formattedDateString
}
