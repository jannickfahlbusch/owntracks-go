//go:generate stringer -type=BatteryStatus -linecomment -trimprefix BatteryStatus

package types

import "time"

// ListResponse is a generic response containing a list of values
type ListResponse struct {
	Results []string
}

type BatteryStatus int

const (
	BatteryStatusUnknown BatteryStatus = iota
	BatteryStatusUnplugged
	BatteryStatusCharging
	BatteryStatusFull
)

type Trigger string

const (
	TriggerPing                        Trigger = "p"
	TriggerCircularRegion              Trigger = "c"
	TriggerBeaconRegion                Trigger = "b"
	TriggerResponse                    Trigger = "r"
	TriggerManual                      Trigger = "u"
	TriggerTimer                       Trigger = "t"
	TriggerFrequentLocationsMonitoring Trigger = "v"
)

type ConnectivityStatus string

const (
	ConnectivityStatusWiFi    ConnectivityStatus = "w"
	ConnectivityStatusOffline ConnectivityStatus = "o"
	ConnectivityStatusMobile  ConnectivityStatus = "m"
)

type Battery struct {
	BatteryLevel  int           `json:"batt"`
	BatteryStatus BatteryStatus `json:"bs"`
}

type WiFi struct {
	SSID  string `json:"SSID"`
	BSSID string `json:"BSSID"`
}

type Location struct {
	EpochTime int64     `json:"tst"`
	Timestamp time.Time `json:"isotst"`
	Received  time.Time `json:"isorcv"`

	Accuracy         float64 `json:"acc"`
	VerticalAccuracy int     `json:"vac"`

	Altitude float64 `json:"alt"`

	Battery
	WiFi

	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
	Radius    int     `json:"rad"`
	Trigger   Trigger `json:"t"`
	TrackerID string  `json:"tid"`
	GeoHash   string  `json:"ghash"`

	Velocity           int                `json:"vel"`
	BarometricPressure float64            `json:"p"`
	ConnectivityStatus ConnectivityStatus `json:"conn"`
	Topic              string             `json:"topic"`
	InRegions          []string           `json:"inregions"`

	Address     string `json:"addr,omitempty"`
	Locality    string `json:"locality,omitempty"`
	CountryCode string `json:"cc,omitempty"`

	DistanceTravelled int `json:"dist"`
}

type LocationList struct {
	Count  int
	Data   []Location
	Status int
}

type Version struct {
	Version string `json:"version"`
}
