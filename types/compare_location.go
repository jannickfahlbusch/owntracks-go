package types

func (location *Location) IsSamePlace(other *Location) bool {
	return location.Latitude == other.Latitude &&
		location.Longitude == other.Longitude
}

func (location *Location) IsSameTime(other *Location) bool {
	return location.Timestamp.Equal(*other.Timestamp)
}

func (location *Location) IsSamePlaceAndTime(other *Location) bool {
	return location.IsSameTime(other) && location.IsSamePlace(other)
}
