package engine

import (
	"beeweb/crawler/fetcher"
	"log"
)

type SimpleEngine struct {

}
func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		log.Printf("Fetching %s", r.Url)
		body,err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("%s", err)
			continue
		}
		parseResult := r.ParseFunc(body)
		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}

func (e SimpleEngine) worker(r Request) (ParseResult, error) {
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		return ParseResult{}, err
	}
	return r.ParseFunc(body), nil
}