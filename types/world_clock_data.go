package types

type WorldClockData struct {
	Id          uint64
	Name        string
	Time        int32
	IsPaused    bool
	TimeMarkers []TimeMarkerData
}
