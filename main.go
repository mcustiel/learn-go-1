// GoTest project main.go
package main

import (
	"fmt"
)

func main() {
	checksList := SearcherSet{RegexpGetter{"bla\\b"}, RegexpGetter{"b."}}

	config := Config{"string", checksList}

	fmt.Print(GetValuesFrom("bla bla bla", checksList), config)
}
