// Thank robots for this file that was generated for you at 2019-03-10 15:41:45.156998841 +0300 MSK m=+0.011863066
package load_images

import (
	"net/url"
)

type TypeIn struct {
	Body       *string  `json:"body"`
	Brand      string   `json:"brand"`
	Generation string   `json:"generation"`
	Model      string   `json:"model"`
	Phone      string   `json:"phone"`
	Price      int      `json:"price"`
	RawImages  []string `json:"rawImages"`
	Year       int      `json:"year"`
}

type TypeOut struct {
	Body       *string  `json:"body"`
	Brand      string   `json:"brand"`
	Generation string   `json:"generation"`
	Model      string   `json:"model"`
	Phone      string   `json:"phone"`
	Photos     []string `json:"photos"`
	Price      int      `json:"price"`
	Year       int      `json:"year"`
}

func Handle(input TypeIn) TypeOut {
	photos := make([]string, 0, len(input.RawImages))
	for _, raw := range input.RawImages {
		u, _ := url.Parse(raw)
		u.Host = "storage.org"
		photos = append(photos, u.String())
	}
	return TypeOut{
		Body:       input.Body,
		Brand:      input.Brand,
		Generation: input.Generation,
		Model:      input.Model,
		Phone:      input.Phone,
		Price:      input.Price,
		Photos:     photos,
		Year:       input.Year,
	}
}
