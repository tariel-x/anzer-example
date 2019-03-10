// Thank robots for this file that was generated for you at 2019-03-10 15:42:27.81836707 +0300 MSK m=+0.010221678
package create

import (
	"math/rand"
)

type TypeIn struct {
	Body       *string  `json:"body"`
	Brand      string   `json:"brand"`
	Generation string   `json:"generation"`
	Model      string   `json:"model"`
	Phone      string   `json:"phone"`
	Photos     []string `json:"photos"`
	Price      int      `json:"price"`
	Year       int      `json:"year"`
}

type TypeOut struct {
	Body       *string  `json:"body"`
	Brand      string   `json:"brand"`
	Generation string   `json:"generation"`
	Id         int      `json:"id"`
	Model      string   `json:"model"`
	Phone      string   `json:"phone"`
	Photos     []string `json:"photos"`
	Price      int      `json:"price"`
	Year       int      `json:"year"`
}

func Handle(input TypeIn) TypeOut {
	id := rand.Intn(100500)
	return TypeOut{
		Body:       input.Body,
		Brand:      input.Brand,
		Generation: input.Generation,
		Id:         id,
		Model:      input.Model,
		Phone:      input.Phone,
		Photos:     input.Photos,
		Price:      input.Price,
		Year:       input.Year,
	}
}
