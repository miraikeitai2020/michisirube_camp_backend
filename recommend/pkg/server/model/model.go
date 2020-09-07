package model

type Request struct {
	Emotion  Emotion  `json:"emotion"`
	Location Location `json:"location"`
	Time     int      `json:"time"`
}
type Emotion struct {
	Value int `json:"value"`
}
type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
