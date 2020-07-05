// Thank robots for this file that was generated for you at 2020-07-05 14:58:13.415304434 +0300 MSK m=+0.040382528
package result

type TypeIn struct {
	Left *struct {
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
	} `json:"left"`
	Right *struct {
		Error string `json:"error"`
	} `json:"right"`
}

type TypeOut struct {
	Car *struct {
		Brand string `json:"brand"`
		Model string `json:"model"`
	} `json:"car"`
	Err *string `json:"err"`
	Id  *string `json:"id"`
}

func Handle(input TypeIn) TypeOut {
	return handler(input)
}
