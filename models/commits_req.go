package models

type ContribRequest struct {
	UserOrg     string `json:"userOrg" valid:"alphanum"`
	Repo        string `json:"repo" valid:"alphanum"`
	Author      string `json:"author" valid:"alphanum"`
	DateSince   string `json:"dateSince" valid:"alphanum"`
	DateUntil   string `json:"dateUntil" valid:"alphanum"`
	AccessToken string `json:"accessToken" valid:"alphanum"`
}
