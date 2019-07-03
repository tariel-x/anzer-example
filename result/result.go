package result

func handler(input TypeIn) TypeOut {
	if input.Right != nil {
		return TypeOut{
			Error: &input.Right.Error,
		}
	} else if input.Left != nil {
		return TypeOut{
			Id: &input.Left.Id,
			Car: &struct {
				Brand string `json:"brand"`
				Model string `json:"model"`
			}{
				Brand: input.Left.Car.Brand,
				Model: input.Left.Car.Model,
			},
		}
	}
	err := "no left neither right"
	return TypeOut{
		Error: &err,
	}
}
