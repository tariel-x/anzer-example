// Thank robots for this file that was generated for you at 2019-03-10 15:36:14.079175184 +0300 MSK m=+0.025789081
package transform

import "strings"

type TypeIn struct {
	Body   *string  `json:"body"`
	Brand  string   `json:"brand"`
	Model  string   `json:"model"`
	Phone  string   `json:"phone"`
	Photos []string `json:"photos"`
	Price  float64  `json:"price"`
	Year   int      `json:"year"`
}

type TypeOut struct {
	Body       *string  `json:"body"`
	Brand      string   `json:"brand"`
	Generation string   `json:"generation"`
	Model      string   `json:"model"`
	Phone      string   `json:"phone"`
	Price      int      `json:"price"`
	RawImages  []string `json:"rawImages"`
	Year       int      `json:"year"`
}

func Handle(input TypeIn) TypeOut {
	modelGeneration := strings.Split(input.Model, " ")
	model := ""
	if len(modelGeneration) >= 1 {
		model = modelGeneration[0]
	}
	generation := ""
	if len(modelGeneration) >= 2 {
		generation = modelGeneration[1]
	}
	return TypeOut{
		Body:       input.Body,
		Brand:      input.Brand,
		Generation: generation,
		Model:      model,
		Phone:      input.Phone,
		Price:      int(input.Price),
		RawImages:  input.Photos,
		Year:       input.Year,
	}
}
