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

const LocationType = "location"

type ConnectivityStatus string

const (
	ConnectivityStatusWiFi    ConnectivityStatus = "w"
	ConnectivityStatusOffline ConnectivityStatus = "o"
	ConnectivityStatusMobile  ConnectivityStatus = "m"
)

type Battery struct {
	BatteryLevel  int           `json:"batt,omitempty"`
	BatteryStatus BatteryStatus `json:"bs,omitempty"`
}

type WiFi struct {
	SSID  string `json:"SSID,omitempty"`
	BSSID string `json:"BSSID,omitempty"`
}

type Location struct {
	Type string `json:"_type,omitempty"`

	EpochTime int64      `json:"tst"`
	Timestamp *time.Time `json:"isotst,omitempty"`
	Received  *time.Time `json:"isorcv,omitempty"`

	Accuracy         float64 `json:"acc,omitempty"`
	VerticalAccuracy float64 `json:"vac,omitempty"`

	Altitude float64 `json:"alt,omitempty"`

	Battery
	WiFi

	Latitude         float64 `json:"lat"`
	Longitude        float64 `json:"lon"`
	Radius           float64 `json:"rad,omitempty"`
	Trigger          Trigger `json:"t,omitempty"`
	TrackerID        string  `json:"tid,omitempty"`
	GeoHash          string  `json:"ghash,omitempty"`
	CourseOverGround float64 `json:"cog,omitempty"`

	Velocity           float64            `json:"vel,omitempty"`
	BarometricPressure float64            `json:"p,omitempty"`
	ConnectivityStatus ConnectivityStatus `json:"conn,omitempty"`
	Topic              string             `json:"topic,omitempty"`
	InRegions          []string           `json:"inregions,omitempty"`

	Address     string `json:"addr,omitempty,omitempty"`
	Locality    string `json:"locality,omitempt,omitempty"`
	CountryCode string `json:"cc,omitempty"`

	DistanceTravelled int `json:"dist,omitempty"`
}

type LocationList struct {
	Count  int
	Data   []*Location
	Status int
}

type Version struct {
	Version string `json:"version"`
}
