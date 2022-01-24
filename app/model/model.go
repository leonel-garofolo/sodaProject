package model

type Client struct {
	Id           int
	Order        int
	PricePerSoda float64
	PricePerBox  float64
	Address      string
	NumAddress   int
	IdDelivery   int
	IdRoot       int
	Due          int
}

type Delivery struct {
	Id   int
	Name string
	Code []int
}

type DeliveryRoot struct {
	IdDelivery int
	IdRoot     int
	Code       int
}
