package main

import (
	/* "modernc.org/tk" */
	"fmt"
	"github.com/takusuman/tkli/lib"
)

const (
	/* It would be better to do this in a configuration file */
	weatherAPI = "http://wttr.in/?format=j1"
)

func main() {
	apiData, _ := lib.FetchAPIData(weatherAPI)
	weatherReport, _ := lib.Data2Struct(apiData)

	fmt.Println(weatherReport.CurrentCondition[0])
}
