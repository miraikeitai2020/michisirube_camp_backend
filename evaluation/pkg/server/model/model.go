package model

type Evaluation struct{
	Value int `json:"value"`
	Emotion Emotion `json:"emotion"`
}
type Emotion struct{
	Value int `json:"value"`
}

type Location struct{
	ID int `json:"id"`
}