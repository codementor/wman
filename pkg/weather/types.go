package weather

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
