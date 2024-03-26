package entities

type Load struct {
	Client        string `json:"client"`
	Plate         string `json:"plate"`
	Material      string `json:"material"`
	Quantity      string `json:"quantity"`
	PaymentMethod string `json:"paymentMethod"`
	Signature     string `json:"signature"`
}

type Client struct {
	Name  string `json:"name"`
	Plate string `json:"plate"`
}

type Signature struct {
	Url string `json:"url"`
}
