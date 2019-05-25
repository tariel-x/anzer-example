// Thank robots for this file that was generated for you at 2019-05-25 16:08:51.973326447 +0300 MSK m=+0.008347686
package create

type EitherRight struct {
	Error string `json:"error"`
}
type HandlerIn struct {
	Body       *string  `json:"body"`
	Brand      string   `json:"brand"`
	Generation string   `json:"generation"`
	Model      string   `json:"model"`
	Phone      string   `json:"phone"`
	Photos     []string `json:"photos"`
	Price      int      `json:"price"`
	Year       int      `json:"year"`
}
type HandlerOut struct {
	Left  *EitherLeft  `json:"left"`
	Right *EitherRight `json:"right"`
}
type EitherLeft struct {
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
}

func Handle(input HandlerIn) HandlerOut {
	return handler(input)
}
