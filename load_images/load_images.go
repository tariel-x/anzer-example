package load_images

import (
	"net/url"
)

func handler(input TypeIn) TypeOut {
	photos := make([]string, 0, len(input.RawImages))
	for _, raw := range input.RawImages {
		u, _ := url.Parse(raw)
		u.Host = "storage.org"
		photos = append(photos, u.String())
	}
	return TypeOut{
		Left: &struct {
			Body       *string  `json:"body"`
			Brand      string   `json:"brand"`
			Generation string   `json:"generation"`
			Model      string   `json:"model"`
			Phone      string   `json:"phone"`
			Photos     []string `json:"photos"`
			Price      int      `json:"price"`
			Year       int      `json:"year"`
		}{
			Body:       input.Body,
			Brand:      input.Brand,
			Generation: input.Generation,
			Model:      input.Model,
			Phone:      input.Phone,
			Price:      input.Price,
			Photos:     photos,
			Year:       input.Year,
		},
		Right: nil,
	}
}
