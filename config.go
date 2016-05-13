package main

type SearcherSet []Searcher

type Config struct {
	startSite string
	checks    SearcherSet
}

func getValuesFrom(input string, checks SearcherSet) map[string]string {
	for i := 0; i < len(checks); i++ {

	}
}
