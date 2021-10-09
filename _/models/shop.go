package models

type Shop struct {
	Customers []Customer
	Products  []Product
}

func (shop *Shop) AddCustomer(email, phone, address string) {
	c := Customer{
		Email:   email,
		Phone:   phone,
		Address: address,
	}
	shop.Customers = append(shop.Customers, c)
}

func (shop *Shop) AddProduct(title string, stQty int, price float64) {
	p := Product{
		Title:         title,
		StockQuantity: stQty,
		Price:         price,
	}
	shop.Products = append(shop.Products, p)
}

func BuyProduct(customer Customer, product Product) {

}
