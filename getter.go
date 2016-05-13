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
	find(input string) (string, bool)
}

func (getter *RegexpGetter) find(input string) (string, bool) {
	var regexpId = regexp.MustCompile(getter.pattern)
	if regexpId.MatchString(input) {
		return regexpId.FindString(input), true
	}
	return "", false
}

func (getter *XpathGetter) find(input string) (string, bool) {
	path := xmlpath.MustCompile(getter.xpath)
	root, err := xmlpath.ParseString(input)
	if err != nil {
		return "", false
	}
	return path.String(root)
}
