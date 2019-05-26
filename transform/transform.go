package transform

import "strings"

func handler(input TypeIn) TypeOut {
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
		Left: &struct {
			Body       *string  `json:"body"`
			Brand      string   `json:"brand"`
			Generation string   `json:"generation"`
			Model      string   `json:"model"`
			Phone      string   `json:"phone"`
			Price      int      `json:"price"`
			RawImages  []string `json:"rawImages"`
			Year       int      `json:"year"`
		}{
			Body:       input.Body,
			Brand:      input.Brand,
			Generation: generation,
			Model:      model,
			Phone:      input.Phone,
			Price:      int(input.Price),
			RawImages:  input.Photos,
			Year:       input.Year,
		},
		Right: nil,
	}
}
