package meteorology

import "fmt"

const (
	Celsius TemperatureUnit = iota
	Fahrenheit
)
const (
	KmPerHour SpeedUnit = iota
	MilesPerHour
)

type Stringer interface {
	String() string
}
type TemperatureUnit int
type SpeedUnit int

type Speed struct {
	magnitude int
	unit      SpeedUnit
}
type Temperature struct {
	degree int
	unit   TemperatureUnit
}
type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

func (st TemperatureUnit) String() string {
	switch st {
	case Celsius:
		return "°C"
	case Fahrenheit:
		return "°F"
	default:
		return fmt.Sprintf("unknown unit (%d)", st)
	}
}
func (t Temperature) String() string {
	return fmt.Sprintf("%v %v", t.degree, t.unit)
}

var speedUnits = map[SpeedUnit]string{
	KmPerHour:    "km/h",
	MilesPerHour: "mph",
}

func (su SpeedUnit) String() string {
	if str, ok := speedUnits[su]; ok {
		return str
	}
	return fmt.Sprintf("unknown unit (%d)", su)
}

func (s Speed) String() string {
	return fmt.Sprintf("%d %v", s.magnitude, s.unit)
}

func (m MeteorologyData) String() string {
	return fmt.Sprintf(
		"%v: %v, Wind %v at %v, %v%% Humidity",
		m.location, m.temperature, m.windDirection,
		m.windSpeed, m.humidity)
}
