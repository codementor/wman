package weather

type UnitType string

const (
	Fahrenheit UnitType = "fahrenheit"
	celsius             = "celsius"
)

type WeatherConfig struct {
	Key    string   `json:"key"`
	Unit   UnitType `json:"unit"`
	Cities []string `json:"cities"`
}
