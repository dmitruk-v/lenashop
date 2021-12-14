package models

type CartItem struct {
	Product
	BuyQuantity int
}

type Cart struct {
	CartId     int
	CustomerId int
	Items      []CartItem
}

func NewCart(customerId int) Cart {
	return Cart{
		CustomerId: customerId,
		Items:      make([]CartItem, 0),
	}
}

func (cart *Cart) Update(updCart Cart) Cart {
	items := append(make([]CartItem, 0), updCart.Items...)
	return Cart{
		CartId:     updCart.CartId,
		CustomerId: updCart.CustomerId,
		Items:      items,
	}
}

func (cart *Cart) Clear() Cart {
	return Cart{
		CartId:     cart.CartId,
		CustomerId: cart.CustomerId,
		Items:      make([]CartItem, 0),
	}
}
