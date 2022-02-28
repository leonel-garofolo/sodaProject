package model

import (
	"github.com/go-playground/validator/v10"
)

type Client struct {
	Id           int     `json:"id"`
	Order        int     `json:"order" validate:"required,min=1,max=4"`
	Address      string  `json:"address" validate:"required,min=3,max=100"`
	NumAddress   int     `json:"num_address"  validate:"required,min=1,max=6"`
	PricePerSoda float64 `json:"price_per_soda"  validate:"required,min=1,max=10"`
	PricePerBox  float64 `json:"price_per_box" validate:"required,min=1,max=10"`
	Debt         float64 `json:"debt" validate:"required,min=1,max=10"`
	IdDelivery   int     `json:"id_delivery"`
	IdRoot       int     `json:"id_root" validate:"required,min=1,max=2"`
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

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

var validate = validator.New()

func ValidateStruct(client Client) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
