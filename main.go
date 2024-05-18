package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type Vreme struct {
	Location struct {
		NameCity string `json:"name"`
		Country  string `json:"country"`
	} `json:"location"`
	Current struct {
		TempC           float32 `json:"temp_c"`
		LastUpdatedTemp string  `json:"last_updated"`
		Condition       struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`
	Forecast struct {
		ForecastDay []struct {
			Hour []struct {
				TimeEpoch int     `json:"time_epoch"`
				TempC     float32 `json:"temp_c"`
				Condition struct {
					Text string `json:"text"`
				} `json:"condition"`
			} `json:"hour"`
			//
			Day struct {
				MaxTempC        float32 `json:"maxtemp_c"`
				DailyChanceRain int     `json:"daily_chance_of_rain"`
			} `json:"day"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

func main() {

	// weatherapi.com personal API key must be provided
	weatherApiKey := ""
	weatherApiCity := "London" //by default
	weatherApiString := "http://api.weatherapi.com/v1/forecast.json?key=" + weatherApiKey + "&q=" + weatherApiCity + "&days=1&aqi=no&alerts=no"
	res, err := http.Get(weatherApiString)

	// if there's an error
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	// check if status code is not 200, which is the one for "OK"
	if res.StatusCode != 200 {
		panic("API not available")
	}

	body, err := io.ReadAll(res.Body)
	// if there's an error
	if err != nil {
		panic(err)
	}

	//fmt.Println(string(body))

	var weather Vreme
	err = json.Unmarshal(body, &weather)
	if err != nil {
		panic(err)
	}

	locationNameCity := weather.Location.NameCity
	lastUpdatedTemp := strings.Split(weather.Current.LastUpdatedTemp, " ")
	currentTemp := weather.Current.TempC
	_ = lastUpdatedTemp[0]                     //lastUpdatedTemp_Date
	lastUpdatedTemp_Hour := lastUpdatedTemp[1] //lastUpdatedTemp_Hour

	hours := weather.Forecast.ForecastDay[0].Hour

	fmt.Println(locationNameCity, " - Last weather update:", lastUpdatedTemp_Hour)
	fmt.Println("Current Temperate: ", currentTemp, "degrees C")

	now := time.Now()
	hourr := now.Hour()
	//fmt.Println(hourr)

	for index, hour := range hours {
		if index < hourr {
			continue
		}
		fmt.Println("At hour", index, "-", hour.TempC, "degrees C")
		//fmt.Printf("%T\n", hour)
	}

}
