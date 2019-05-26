package create

import (
	"math/rand"
)

func handler(input TypeIn) TypeOut {
	id := rand.Intn(100500)
	return TypeOut{
		Left: &struct {
			Body *string `json:"body"`
			Car  struct {
				Brand string `json:"brand"`
				Model string `json:"model"`
			} `json:"car"`
			Generation string   `json:"generation"`
			Id         int      `json:"id"`
			Phone      string   `json:"phone"`
			Photos     []string `json:"photos"`
			Price      int      `json:"price"`
			Year       int      `json:"year"`
		}{
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
		},
		Right: nil,
	}
}
