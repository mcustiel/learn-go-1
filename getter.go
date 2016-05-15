package main

import (
	"regexp"
	"strings"

	"launchpad.net/xmlpath"
)

type RegexpGetter struct {
	pattern string
}

type RegexpSubmatchGetter struct {
	pattern string
}

type XpathGetter struct {
	xpath string
}

type Searcher interface {
	Find(input string) ([]string, bool)
}

func (getter RegexpGetter) Find(input string) ([]string, bool) {
	regexpId := regexp.MustCompile(getter.pattern)
	if regexpId.MatchString(input) {
		return regexpId.FindAllString(input, -1), true
	}
	return nil, false
}

func (getter RegexpSubmatchGetter) Find(input string) ([]string, bool) {
	if regexpId := regexp.MustCompile(getter.pattern); regexpId.MatchString(input) {
		response := regexpId.FindAllStringSubmatch(input, -1)
		result := make([]string, len(response))
		for index := 0; index < len(response); index++ {
			result[index] = response[index][1]
		}
		return result, true
	}
	return nil, false
}

func (getter XpathGetter) Find(input string) ([]string, bool) {
	path := xmlpath.MustCompile(getter.xpath)
	root, err := xmlpath.Parse(strings.NewReader(input))
	if err != nil {
		return nil, false
	}
	return getFromIterator(path.Iter(root)), true
}

func getFromIterator(iterator *xmlpath.Iter) []string {
	response := make([]string, 0)
	index := 0
	for iterator.Next() {
		response[index] = iterator.Node().String()
		index++
	}
	return response
}
