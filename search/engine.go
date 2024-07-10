package search

import (
	"animaya/search-engine/db"
	"fmt"
	"time"
)

func RunEngine() {
	fmt.Println("started search engine crawl...")
	defer fmt.Println("search engine crawl has finished")

	settings := &db.SearchSettings{}

	err := settings.Get()

	if err != nil {

		fmt.Println("something went wrong getting the settings")
		return
	}

	if !settings.SearchOn {
		fmt.Println("search is turned of")
		return
	}

	crawl := &db.CrawledUrl{}

	nextUrls, err := crawl.GetNextCrawlUrls(int(settings.Amount))

	if err != nil {
		fmt.Println("something went wrong getting next")
		return
	}

	newUrls := []db.CrawledUrl{}

	testedTime := time.Now()

	for _, next := range nextUrls {
		result := runCrawl(next.Url)

		if !result.Success {
			err := next.UpdateUrl(db.CrawledUrl{
				ID:              next.ID,
				Url:             next.Url,
				Success:         false,
				CrawlDuration:   result.CarwlData.CrawlTime,
				ResponseCode:    result.ResponseCode,
				PageTitle:       result.CarwlData.PageTitle,
				PageDescription: result.CarwlData.PageDescription,
				Heading:         result.CarwlData.Headings,
				LastTested:      &testedTime,
			})

			if err != nil {
				fmt.Println("something went wrong updating a failed url")
			}
			continue
		}
		err := next.UpdateUrl(db.CrawledUrl{
			ID:              next.ID,
			Url:             next.Url,
			Success:         result.Success,
			CrawlDuration:   result.CarwlData.CrawlTime,
			ResponseCode:    result.ResponseCode,
			PageTitle:       result.CarwlData.PageTitle,
			PageDescription: result.CarwlData.PageDescription,
			Heading:         result.CarwlData.Headings,
			LastTested:      &testedTime,
		})
		if err != nil {
			fmt.Println("something went wrong updating a success url")
			fmt.Println(next.Url)
		}

		for _, newUrl := range result.CarwlData.Links.External {
			newUrls = append(newUrls, db.CrawledUrl{Url: newUrl})
		}
	}

	if !settings.AddNew {
		return
	}

	for _, newUrl := range newUrls {
		err := newUrl.Save()

		if err != nil {
			fmt.Println("something went wrong adding the new urlto the databse")
		}
	}

	fmt.Printf("\n Added %d new urls to the database", len(newUrls))
}

func RunIndex() {
	fmt.Println("started search indexing...")
	defer fmt.Println("search indexing has finished")

	crawled := &db.CrawledUrl{}

	notIndexed, err := crawled.GetNotIndex()

	if err != nil {
		return
	}
	idx := make(Index)

	idx.Add(notIndexed)

	searchIndex := &db.SearchIndex{}
	err = searchIndex.Save(idx, notIndexed)
	if err != nil {
		return
	}

	err = crawled.SetIndexedTrue(notIndexed)
	if err != nil {
		return
	}

}
