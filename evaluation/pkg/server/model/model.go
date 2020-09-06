package model
type Request struct{
	Evaluation Evaluation `json:"Evaluation"`
	Location	Location	`json::"Location"`
}
type Evaluation struct{
	Value int `json:"value"`
	Emotion Emotion `json:"emotion"`
}
type Emotion struct{
	Value int `json:"value"`
}

type Location struct{
	Id int `json:"id"`
}