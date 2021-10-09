package models

type OrderItem struct {
	Product
	BuyQuantity int
}

type Order struct {
	OrderId int
	Items   []OrderItem
}

type Customer struct {
	CustomerId int
	Email      string
	Phone      string
	Address    string
	Cart       Cart
	Orders     []Order
}

func CreateCustomer(email string, phone string, address string) Customer {
	return Customer{
		Email:   email,
		Phone:   phone,
		Address: address,
	}
}

func (customer *Customer) AddOrder(cart Cart) {
	var order Order
	for _, cit := range cart.Items {
		order.Items = append(order.Items, OrderItem(cit))
	}
	customer.Orders = append(customer.Orders, order)
	customer.Cart.Clear()
}

func (customer *Customer) AddToCart(product Product) {
	customer.Cart.AddItem(product, 1)
}
