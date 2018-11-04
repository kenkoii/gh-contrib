package models

type ContribRequest struct {
	UserOrg   string `json:"userOrg" valid:"ascii"`
	Repo      string `json:"repo" valid:"ascii"`
	Author    string `json:"author" valid:"ascii"`
	DateSince string `json:"dateSince" valid:"ascii"`
	DateUntil string `json:"dateUntil" valid:"ascii"`
	// AccessToken string `json:"accessToken" valid:"alphanum"`
}
