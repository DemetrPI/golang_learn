package meteorology

import "fmt"

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)
const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
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
	units := []string{"°C", "°F"}
	return units[st]
}
func (t Temperature) String() string {
	return fmt.Sprintf("%v %v", t.degree, t.unit)
}

func (su SpeedUnit) String() string {
	units := []string{"km/h", "mph"}
	return units[su]
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
