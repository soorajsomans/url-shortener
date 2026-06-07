package dto

type ShortenRequest struct {
	URL string `json: "url" example:"https://google.com"`
}
