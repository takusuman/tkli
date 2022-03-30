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
	/* Expect getting a "struct array" back */
	weatherReport, _ := lib.Data2Struct(apiData)
	weatherCurrentCondition := weatherReport.CurrentCondition[0]
	weatherArea := weatherReport.NearestArea[0]

	fmt.Printf("TempC: %s\nArea: %s\n", weatherCurrentCondition.TempC, weatherArea.AreaName[0].Value)
}
