package models

type UrlProtected struct {
	Id    uint64   `json:"id"`
	Url   string   `json:"url"`
	Roles []string `json:"roles"`
}
