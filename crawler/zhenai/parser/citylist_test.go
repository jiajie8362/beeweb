package parser

import (
	"testing"
	"beeweb/crawler/fetcher"
)

func TestParseCityList(t *testing.T) {
	contents, err := fetcher.Fetch("http://www.zhenai.com/zhenhu")
	if err != nil {
		panic(err)
	}
	result := ParseCityList(contents)
	const resultSize = 470
	if len(result.Requests) != resultSize {
		t.Errorf("result should have 470")
	}
}
