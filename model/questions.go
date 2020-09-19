package model

type Question struct {
	Name        string
	Description string
	Type        string
	Example     string
}

type Storekeeper struct {
	Question
}
