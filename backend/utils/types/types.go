package types

type Record struct {
	Name      string `json:"name"`
	Descr     string `json:"descr"`
	Image     string `json:"img"`
	Price     string `json:"price"`
	Topic     string `json:"topic"`
	Created   string `json:"created"`
	Published string `json:"published"`
	Quant     string `json:"quantity"`
}

type Response struct {
	Status int      `json:"status"`
	Msg    string   `json:"msg"`
	Data   []Record `json:"data"`
}
