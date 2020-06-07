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
