package main

import (
	/* "modernc.org/tk" */
	"fmt"
)

const (
	/* It would be better to do this in a configuration file */
	weatherAPI = "http://wttr.in/?format=j1"
)

func main() {
	apiData, _, _ := fetch(weatherAPI)
	weatherReport := data2struct(apiData)

	fmt.Println(weatherReport.CurrentCondition[0])
}
