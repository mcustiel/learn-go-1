// GoTest project main.go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var linkChecks SearcherSet

func main() {
	checksList := SearcherSet{"title": RegexpGetter{"<h4>[^<]+</h4>"}, "subtitle": RegexpGetter{"<h1\\s+itemprop=\"headline\">[^<]+</h1>"}}

	linkChecks = SearcherSet{"link": RegexpGetter{"https?://[!\"]+"}}

	visited := make([]string, 0)
	found := make(map[string]Coincidences, 0)

	fmt.Println(crawl(
		"http://www.montevideo.com.uy/auc.aspx?308257,3",
		checksList,
		0,
		1,
		&visited,
		found))
}

func crawl(
	site string,
	checksList SearcherSet,
	level int,
	maxLevel int,
	visited *[]string,
	found map[string]Coincidences) map[string]Coincidences {
	response, err := http.Get(site)
	if err != nil {
		log.Fatal(err)
	}
	content, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	found[site] = GetValuesFrom(string(content), checksList)

	return found
}
