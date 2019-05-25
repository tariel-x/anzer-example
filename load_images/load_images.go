// Thank robots for this file that was generated for you at 2019-03-10 15:41:45.156998841 +0300 MSK m=+0.011863066
package load_images

import (
	"net/url"
)

func handler(input HandlerIn) HandlerOut {
	photos := make([]string, 0, len(input.RawImages))
	for _, raw := range input.RawImages {
		u, _ := url.Parse(raw)
		u.Host = "storage.org"
		photos = append(photos, u.String())
	}
	model := EitherLeft{
		Body:       input.Body,
		Brand:      input.Brand,
		Generation: input.Generation,
		Model:      input.Model,
		Phone:      input.Phone,
		Price:      input.Price,
		Photos:     photos,
		Year:       input.Year,
	}
	return HandlerOut{
		Left:  &model,
		Right: nil,
	}
}
