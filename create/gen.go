// Thank robots for this file that was generated for you at 2019-05-26 23:08:30.970732446 +0300 MSK m=+0.008845662
package create

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
	Left *struct {
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
	} `json:"left"`
	Right *struct {
		Error string `json:"error"`
	} `json:"right"`
}

func Handle(input TypeIn) TypeOut {
	return handler(input)
}
