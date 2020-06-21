package create

import (
	"context"
	"encoding/json"
	"os"

	_ "github.com/go-kivik/couchdb/v4"
	kivik "github.com/go-kivik/kivik/v4"
)

type DbModel struct {
	OutLeft
	Id string `json:"-"`
}

func handler(input TypeIn) TypeOut {
	client, err := kivik.New("couch", os.Getenv("CLOUDANT_URL"))
	if err != nil {
		return right(err)
	}

	db := client.DB(context.TODO(), "posts")

	out := OutLeft{
		Body: input.Body,
		Car: OutLeftCar{
			Brand: input.Brand,
			Model: input.Model,
		},
		Generation: input.Generation,
		Phone:      input.Phone,
		Photos:     input.Photos,
		Price:      input.Price,
		Year:       input.Year,
	}

	newDoc := DbModel{
		OutLeft: out,
	}

	doc, err := json.Marshal(newDoc)
	if err != nil {
		return right(err)
	}

	id, _, err := db.CreateDoc(context.TODO(), doc)
	if err != nil {
		return right(err)
	}

	out.Id = id

	return Left(out)
}

func right(err error) TypeOut {
	return TypeOut{
		Right: &struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		},
	}
}

// MUST BE GENERATED BELOW

type OutLeft struct {
	Body       *string    `json:"body"`
	Car        OutLeftCar `json:"car"`
	Generation string     `json:"generation"`
	Id         string     `json:"id"`
	Phone      string     `json:"phone"`
	Photos     []string   `json:"photos"`
	Price      int        `json:"price"`
	Year       int        `json:"year"`
}

type OutLeftCar struct {
	Brand string `json:"brand"`
	Model string `json:"model"`
}

type OutRight struct {
	Error string `json:"error"`
}

func Left(input OutLeft) TypeOut {
	return TypeOut{
		Left: &struct {
			Body *string `json:"body"`
			Car  struct {
				Brand string `json:"brand"`
				Model string `json:"model"`
			} `json:"car"`
			Generation string   `json:"generation"`
			Id         string   `json:"id"`
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
				Brand: input.Car.Brand,
				Model: input.Car.Model,
			},
			Generation: input.Generation,
			Id:         input.Id,
			Phone:      input.Phone,
			Photos:     input.Photos,
			Price:      input.Price,
			Year:       input.Year,
		},
		Right: nil,
	}
}

func Right(input OutRight) TypeOut {
	return TypeOut{
		Right: &struct {
			Error string `json:"error"`
		}{
			Error: input.Error,
		},
	}
}
