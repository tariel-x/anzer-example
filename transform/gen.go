// Thank robots for this file that was generated for you at 2019-05-26 23:06:23.414865258 +0300 MSK m=+0.010342797
package transform

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
	Left *struct {
		Body       *string  `json:"body"`
		Brand      string   `json:"brand"`
		Generation string   `json:"generation"`
		Model      string   `json:"model"`
		Phone      string   `json:"phone"`
		Price      int      `json:"price"`
		RawImages  []string `json:"rawImages"`
		Year       int      `json:"year"`
	} `json:"left"`
	Right *struct {
		Error string `json:"error"`
	} `json:"right"`
}

func Handle(input TypeIn) TypeOut {
	return handler(input)
}
