// Thank robots for this file that was generated for you at 2019-03-10 15:36:14.079175184 +0300 MSK m=+0.025789081
package transform

import "strings"

func handler(input HandlerIn) HandlerOut {
	modelGeneration := strings.Split(input.Model, " ")
	model := ""
	if len(modelGeneration) >= 1 {
		model = modelGeneration[0]
	}
	generation := ""
	if len(modelGeneration) >= 2 {
		generation = modelGeneration[1]
	}
	ret := EitherLeft{
		Body:       input.Body,
		Brand:      input.Brand,
		Generation: generation,
		Model:      model,
		Phone:      input.Phone,
		Price:      int(input.Price),
		RawImages:  input.Photos,
		Year:       input.Year,
	}
	return HandlerOut{
		Left:  &ret,
		Right: nil,
	}
}
