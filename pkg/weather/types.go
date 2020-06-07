package weather

import (
	"sort"
)

type UnitType string

const (
	Fahrenheit UnitType = "fahrenheit"
	Celsius    UnitType = "celsius"
)

type Config struct {
	Key    string   `json:"key"`
	Unit   UnitType `json:"unit"`
	Cities []string `json:"cities"`
}

// Structures required for open weather map structures
// https://openweathermap.org/current

type main struct {
	Temp float64 `json:"temp"`
}

type description struct {
	ID   int    `json:"id"`
	Desc string `json:"description"`
}

type weather struct {
	Name     string        `json:"name"`
	Main     main          `json:"main"`
	Weathers []description `json:"weather"`
}

// App weather structure
type Model struct {
	City string
	Temp float64
	Desc string
}

type Models []Model

func (m Models) Len() int {
	return len(m)
}

func (m Models) Less(i, j int) bool {
	// what do you want to sort on?  city?  temperature?
	return m[i].City < m[j].City
}

func (m Models) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

var _ sort.Interface = &Models{}
