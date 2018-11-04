package models

type DailyContrib struct {
	Date         string `json:"date"`
	Author       string `json:"author"`
	CommitsCount int    `json:"commitsCount"`
}
