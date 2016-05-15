// GoTest project main.go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

var linkChecks SearcherSet

func main() {
	checksList := SearcherSet{"title": RegexpSubmatchGetter{"<h4>([^<]+)</h4>"}, "subtitle": RegexpSubmatchGetter{"<h1\\b.*itemprop=\"headline\"[^>]*>([^<]+)</h1>"}}

	linkChecks = SearcherSet{"link": RegexpSubmatchGetter{"<a\\b.*href=\"([^\"]+)\"[^>]*>"}}

	visited := make(map[string]bool, 0)
	found := make(map[string]Coincidences, 0)

	crawl(
		"http://www.montevideo.com.uy/auc.aspx?308257,3",
		checksList,
		0,
		3,
		visited,
		found)
	for site, data := range found {
		fmt.Println("* Site: ", site)
		for property, values := range data {
			fmt.Println("** Property: ", property)
			for i := 0; i < len(values); i++ {
				fmt.Println("*** ", values[i])
			}
		}
		fmt.Println("")
	}
	fmt.Println(found)
}

func crawl(
	site string,
	checksList SearcherSet,
	level int,
	maxLevel int,
	visited map[string]bool,
	found map[string]Coincidences) {
	if level <= maxLevel {
		response, err := http.Get(site)
		if err != nil {
			log.Fatal(err)
		}
		content, err := ioutil.ReadAll(response.Body)
		response.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
		foundValues := GetValuesFrom(string(content), checksList)
		_, hasTitle := foundValues["title"]
		_, hasSubtitle := foundValues["subtitle"]
		if hasTitle && hasSubtitle {
			found[site] = foundValues
		}

		visited[sanitizeUrl(site)] = true
		links := GetValuesFrom(string(content), linkChecks)
		for i := 0; i < len(links["link"]); i++ {
			if uri, _ := url.Parse(links["link"][i]); strings.Contains(uri.Scheme, "http") && strings.HasPrefix(uri.Host, "www.montevideo.com.uy") {
				if _, exists := visited[sanitizeUrl(links["link"][i])]; !exists {
					crawl(
						links["link"][i],
						checksList,
						level+1,
						maxLevel,
						visited,
						found)
				}
			}
		}
	}
}

func sanitizeUrl(url string) string {
	return strings.Replace(strings.Replace(url, "https://", "", 1), "http://", "", 1)
}
