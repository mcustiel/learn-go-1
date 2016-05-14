package main

type SearcherSet []Searcher

type Config struct {
	startSite string
	checks    SearcherSet
}

func GetValuesFrom(input string, checks SearcherSet) []string {
	found := make([]string, 0)
	for i := 0; i < len(checks); i++ {
		result, ok := checks[i].Find(input)
		if ok {
			for j := 0; j < len(result); j++ {
				found = append(found, result[i])
			}
		}
	}
	return found
}
