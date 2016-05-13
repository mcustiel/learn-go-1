package main

import "regexp"
import "launchpad.net/xmlpath"

type RegexpGetter struct {
	pattern string
}

type XpathGetter struct {
	xpath string
}

type Searcher interface {
	find(input string) ([]string, bool)
}

func (getter *RegexpGetter) find(input string) ([]string, bool) {
	regexpId := regexp.MustCompile(getter.pattern)
	if regexpId.MatchString(input) {
		return regexpId.FindAllString(input, -1), true
	}
	return nil, false
}

func (getter *XpathGetter) find(input string) ([]string, bool) {
	path := xmlpath.MustCompile(getter.xpath)
	root, err := xmlpath.ParseString(input)
	if err != nil {
		return nil, false
	}
	var response []string = {}
	 
	iterator := path.Iter(root)
	for ;iterator.Next(); {
		response[] = iterator.Node()
	}
	return response, true
}
