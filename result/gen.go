// Thank robots for this file that was generated for you at 2019-07-03 23:30:54.635188718 +0300 MSK m=+0.042866705
package result

type TypeIn struct {
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

type TypeOut struct {
	Car *struct {
		Brand string `json:"brand"`
		Model string `json:"model"`
	} `json:"car"`
	Error *string `json:"error"`
	Id    *int    `json:"id"`
}

func Handle(input TypeIn) TypeOut {
	return handler(input)
}
