package main
// Basic response struct
type BasicResponse struct {
    Code       int `json:"code"`
    Message     string `json:"message"`
}

type repositorySummary struct {
	ID         int
	Name       string
	Owner      string
	TotalStars int
}

type repositories struct {
	Repositories []repositorySummary `json:"res"`
}

