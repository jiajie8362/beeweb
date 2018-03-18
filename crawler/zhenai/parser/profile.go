package parser

import (
	"beeweb/crawler/engine"
	"regexp"
	"strconv"
	"beeweb/model"
)

const ageRe = `<td><span class="label">年龄:</span>([\d]+)岁</td>`

func ParseProfile(contents []byte, name string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name
	age, err := strconv.Atoi(extractString(contents, regexp.MustCompile(ageRe)))
	if err != nil {
		profile.Age = age
	}
	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	}
	return ""
}
