package model

type Client struct {
	Id           int     `json:"id"`
	Order        int     `json:"order"`
	PricePerSoda float64 `json:"price_per_soda"`
	PricePerBox  float64 `json:"price_per_box"`
	Address      string  `json:"address"`
	NumAddress   int     `json:"num_address"`
	IdDelivery   int     `json:"id_delivery"`
	IdRoot       int     `json:"id_root"`
	Due          int     `json:"due"`
}

type Delivery struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Code []int
}

type DeliveryRoot struct {
	IdDelivery int `json:"id_delivery"`
	IdRoot     int `json:"id_root"`
	Code       int `json:"code"`
}
