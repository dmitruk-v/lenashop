package models

type CartItem struct {
	Product
	BuyQuantity int
}

type Cart struct {
	CartId int
	Items  []CartItem
}

func (cart *Cart) CreateCartItem(p Product, bq int) CartItem {
	return CartItem{Product: p, BuyQuantity: bq}
}

func (cart *Cart) AddItem(product Product, buyQuantity int) {
	cart.Items = append(cart.Items, cart.CreateCartItem(product, buyQuantity))
}

func (cart *Cart) RemoveItem(productId int) {
	var result []CartItem
	for _, item := range cart.Items {
		if productId != item.ProductId {
			result = append(result, item)
		}
	}
	cart.Items = result
}

func (cart *Cart) UpdateItem(productId int, buyQuantity int) {
	for _, item := range cart.Items {
		if productId == item.ProductId {
			item.BuyQuantity = buyQuantity
		}
	}
}

func (cart *Cart) Clear() {
	cart.Items = nil
}
