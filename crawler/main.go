package main

import (
	"beeweb/crawler/engine"
	"beeweb/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenhun",
		ParseFunc: parser.ParseCityList,
	})
}


