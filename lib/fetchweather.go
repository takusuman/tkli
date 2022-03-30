package lib 

import (
	"encoding/json"
	"io"
	"net/http"
)

/* Generated automatically by JSON-to-Go, because I do not hate myself for doing
* this by hand.
 */

type weatherReport_t struct {
	CurrentCondition []struct {
		FeelsLikeC       string `json:"FeelsLikeC"`
		FeelsLikeF       string `json:"FeelsLikeF"`
		Cloudcover       string `json:"cloudcover"`
		Humidity         string `json:"humidity"`
		LocalObsDateTime string `json:"localObsDateTime"`
		ObservationTime  string `json:"observation_time"`
		PrecipInches     string `json:"precipInches"`
		PrecipMM         string `json:"precipMM"`
		Pressure         string `json:"pressure"`
		PressureInches   string `json:"pressureInches"`
		TempC            string `json:"temp_C"`
		TempF            string `json:"temp_F"`
		UvIndex          string `json:"uvIndex"`
		Visibility       string `json:"visibility"`
		VisibilityMiles  string `json:"visibilityMiles"`
		WeatherCode      string `json:"weatherCode"`
		WeatherDesc      []struct {
			Value string `json:"value"`
		} `json:"weatherDesc"`
		WeatherIconURL []struct {
			Value string `json:"value"`
		} `json:"weatherIconUrl"`
		Winddir16Point string `json:"winddir16Point"`
		WinddirDegree  string `json:"winddirDegree"`
		WindspeedKmph  string `json:"windspeedKmph"`
		WindspeedMiles string `json:"windspeedMiles"`
	} `json:"current_condition"`
}

func FetchAPIData(apiAddress string) (apiData io.ReadCloser, fetchErr error){
	/* First of all, fetch the JSON that our API (for now, wttr.in) will
	* serve us. */
	response, fetchErr := http.Get(apiAddress)
	
	/* Assign "apiData" to response.Body */
	apiData = response.Body

	return apiData, fetchErr
}

func Data2Struct(apiData io.ReadCloser) (weatherReport weatherReport_t, err error){
	weatherReport = weatherReport_t{}
	err = json.NewDecoder(apiData).Decode(&weatherReport)
	return weatherReport, err
}
