package models

type DailyContrib struct {
	// Date   string `json:"date"`
	Author  string `json:"author"`
	Commits int    `json:"commits"`
}
