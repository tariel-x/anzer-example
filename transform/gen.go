// Thank robots for this file that was generated for you at 2019-05-25 16:36:56.629786978 +0300 MSK m=+0.009918061
package transform

type HandlerIn struct {
	Body   *string  `json:"body"`
	Brand  string   `json:"brand"`
	Model  string   `json:"model"`
	Phone  string   `json:"phone"`
	Photos []string `json:"photos"`
	Price  float64  `json:"price"`
	Year   int      `json:"year"`
}
type HandlerOut struct {
	Left  *EitherLeft  `json:"left"`
	Right *EitherRight `json:"right"`
}
type EitherLeft struct {
	Body       *string  `json:"body"`
	Brand      string   `json:"brand"`
	Generation string   `json:"generation"`
	Model      string   `json:"model"`
	Phone      string   `json:"phone"`
	Price      int      `json:"price"`
	RawImages  []string `json:"rawImages"`
	Year       int      `json:"year"`
}
type EitherRight struct {
	Error string `json:"error"`
}

func Handle(input HandlerIn) HandlerOut {
	return handler(input)
}
