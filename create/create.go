package create

import (
	"math/rand"
)

func handler(input HandlerIn) HandlerOut {
	id := rand.Intn(100500)
	model := EitherLeft{
		Body: input.Body,
		Car: struct {
			Brand string `json:"brand"`
			Model string `json:"model"`
		}{
			Brand: input.Brand,
			Model: input.Model,
		},
		Generation: input.Generation,
		Id:         id,
		Phone:      input.Phone,
		Photos:     input.Photos,
		Price:      input.Price,
		Year:       input.Year,
	}
	return HandlerOut{
		Left:  &model,
		Right: nil,
	}
}
