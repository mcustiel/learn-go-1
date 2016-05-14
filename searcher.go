package main

type SearcherSet map[string]Searcher

type Coincidences map[string][]string

func GetValuesFrom(input string, checks SearcherSet) Coincidences {
	found := make(map[string][]string, 0)
	for id, getter := range checks {
		result, ok := getter.Find(input)
		if ok {
			found[id] = result
		}
	}
	return found
}
