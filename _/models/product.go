package models

import (
	"database/sql"
	"time"
)

type Product struct {
	ProductId     int
	Title         string
	StockQuantity int
	Price         float64
}

type ProductDAO struct {
	ProductId     int
	Title         string
	StockQuantity int
	Price         float64
	CreatedAt     time.Time
	UpdatedAt     sql.NullTime
}

func (p *Product) ToDAO() ProductDAO {
	return ProductDAO{
		ProductId:     p.ProductId,
		Title:         p.Title,
		StockQuantity: p.StockQuantity,
		Price:         p.Price,
	}
}

func (p *Product) FromDAO(pdao ProductDAO) {
	p.ProductId = pdao.ProductId
	p.Title = pdao.Title
	p.StockQuantity = pdao.StockQuantity
	p.Price = pdao.Price
}
