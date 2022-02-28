package services

import "testing"

func TestGetClientForDelivery(t *testing.T) {
	ds := DeliveryService{}
	ds.GetClientForDelivery(nil)

}
