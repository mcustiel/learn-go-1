package main

type SearcherSet []Searcher

type Config struct {
	startSite string
	checks    SearcherSet
}

func GetValuesFrom(input string, checks SearcherSet) []string {
	found := make([]string, 0)
	for i := 0; i < len(checks); i++ {
		result, err := checks[i].Find(input)
		if !err {
			for j := 0; j < len(result); j++ {
				found = append(found, result[i])
			}
		}
	}
	return found
}
