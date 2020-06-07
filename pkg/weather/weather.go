package weather

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

const (
	wurl = "https://api.openweathermap.org/data/2.5/weather"
)

type Fetcher struct {
	url string
}

func New(config *Config) (*Fetcher, error) {

	if strings.TrimSpace(config.Key) == "" {
		return nil, fmt.Errorf("apikey required")
	}
	// open weather defines units as metric or imperial
	var units string
	switch config.Unit {
	case Celsius:
		units = "metrics"
	case Fahrenheit:
		fallthrough
	default:
		units = "imperial"
	}

	return &Fetcher{
		url: fmt.Sprintf("%s?appid=%s&units=%s", wurl, config.Key, units),
	}, nil
}

func (f *Fetcher) Get(city string) (*Model, error) {

	// adding city to the query str
	req, _ := http.NewRequest("GET", f.url, nil)
	req.Header.Add("Accept", "application/json")
	q := req.URL.Query()
	q.Add("q", city)
	req.URL.RawQuery = q.Encode()

	client := http.Client{}
	r, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	switch r.StatusCode {
	case 200:
	case 401:
		return nil, errors.New("not Authorized.  bad or missing key")
	default:
		return nil, fmt.Errorf("unknown HTTP status %d", r.StatusCode)
	}

	weather := new(weather)
	err = json.NewDecoder(r.Body).Decode(weather)
	if err != nil {
		return nil, err
	}

	return convertToModel(weather), nil
}

func convertToModel(w *weather) *Model {
	desc := ""
	if len(w.Weathers) > 0 {
		desc = w.Weathers[0].Desc
	}
	return &Model{
		City: w.Name,
		Temp: w.Main.Temp,
		Desc: desc,
	}
}
