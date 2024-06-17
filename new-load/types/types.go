package types

type Load struct {
	Id            string `json:"id"`
	Client        string `json:"client"`
	Plate         string `json:"plate"`
	Material      string `json:"material"`
	Quantity      string `json:"quantity"`
	PaymentMethod string `json:"paymentMethod"`
	Signature     string `json:"signature"`
}
