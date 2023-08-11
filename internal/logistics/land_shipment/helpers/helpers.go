package helpers

func GetDiscount(productQuantity int, deliveryPrice float64) float64 {
	discount := 0.0
	if productQuantity > 10 {
		discount = deliveryPrice * 0.05
	}

	return discount
}
