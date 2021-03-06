package create

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "github.com/go-kivik/couchdb/v4"
	kivik "github.com/go-kivik/kivik/v4"
)

type DbModel struct {
	OutLeft
	Id string `json:"-"`
}

func handler(input TypeIn) TypeOut {
	log.Println(os.Environ())
	client, err := kivik.New("couch", os.Getenv("__ANZ_CLOUDANT_URL"))
	if err != nil {
		return right(fmt.Errorf("can not connect couch: %q", err))
	}

	db := client.DB(context.Background(), "posts")

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
		return right(fmt.Errorf("can not marshall new doc: %q", err))
	}

	id, _, err := db.CreateDoc(context.Background(), doc)
	if err != nil {
		return right(fmt.Errorf("can not save new doc: %q", err))
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
