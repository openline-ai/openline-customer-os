package model

type Application struct {
	ID          string          `json:"id"`
	Platform    string          `json:"platform"`
	Name        string          `json:"name"`
	TrackerName string          `json:"trackerName"`
	Sessions    AppSessionsPage `json:"sessions"`
	Tenant      string
}
