package products

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
	"regexp"
	"strings"
	"time"

	"dmitrook.ru/lenashop/common"
)

type Product struct {
	ProductId   int
	CategoryId  int
	Title       string
	Price       float64
	Quantity    int
	Description string
	CreatedAt   time.Time
	UpdatedAt   sql.NullTime
}

type ProductImage struct {
	ImageId   int
	ProductId int
	ImageUrl  string
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type FullProduct struct {
	Product
	Images []ProductImage
}

func (p *Product) String() string {
	return fmt.Sprintf("Product{ProductId=%d,CategoryId=%d,Title=%s,Price=%f,Quantity=%d,Description=%s,CreatedAt=%v,UpdatedAt=%v}", p.ProductId, p.CategoryId, p.Title, p.Price, p.Quantity, p.Description, p.CreatedAt, p.UpdatedAt)
}

// var keysReg = regexp.MustCompile(`^(sort|limit)$`)
var sortParamReg = regexp.MustCompile(`^(product_id|category_id|title|price|quantity) (asc|desc)$`)
var limitParamReg = regexp.MustCompile(`^[\d]{1,100}$`)

var defaultOptions = map[string]string{
	"sort":  "product_id ASC",
	"limit": "6",
}

func Products(query url.Values) ([]Product, error) {
	sortQuery := getSortQuery(query["sort"])
	limitQuery := getLimitQuery(query["limit"])

	sqlQuery := fmt.Sprintf("SELECT * FROM product ORDER BY %s, product_id ASC LIMIT %s", sortQuery, limitQuery)
	rows, err := common.DbPool.Query(context.Background(), sqlQuery)
	if err != nil {
		return nil, fmt.Errorf("Products(%v): %v", query, err)
	}
	defer rows.Close()

	var prs []Product
	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ProductId, &p.CategoryId, &p.Title, &p.Price, &p.Quantity, &p.Description, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("Products(%v): %v", query, err)
		}
		prs = append(prs, p)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Products(%v): %v", query, err)
	}
	return prs, nil
}

func ProductsWithImages(query url.Values) ([]FullProduct, error) {
	prs, err := Products(query)
	if err != nil {
		return nil, fmt.Errorf("ProductsWithImages(%v): %v", query, err)
	}
	var fprs []FullProduct
	for _, p := range prs {
		images, err := ImagesByProductId(p.ProductId)
		if err != nil {
			return nil, fmt.Errorf("ProductsWithImages(%v): %v", query, err)
		}
		fprs = append(fprs, FullProduct{p, images})
	}
	return fprs, nil
}

func ImagesByProductId(id int) ([]ProductImage, error) {
	fail := func(err error) error {
		return fmt.Errorf("ImagesByProductId(%v): %v", id, err)
	}
	rows, err := common.DbPool.Query(context.Background(), "SELECT * FROM product_image WHERE product_id = $1", id)
	if err != nil {
		return nil, fail(err)
	}
	var images []ProductImage
	for rows.Next() {
		var img ProductImage
		if err := rows.Scan(&img.ImageId, &img.ProductId, &img.ImageUrl, &img.CreatedAt, &img.UpdatedAt); err != nil {
			return nil, fail(err)
		}
		images = append(images, img)
	}
	return images, nil
}

func ProductById(productId int) (FullProduct, error) {
	row := common.DbPool.QueryRow(context.Background(), "SELECT * FROM product WHERE product_id = $1", productId)
	var fp FullProduct
	var p Product
	err := row.Scan(&p.ProductId, &p.CategoryId, &p.Title, &p.Price, &p.Quantity, &p.Description, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		return fp, fmt.Errorf("OneById(%v): %v", productId, err)
	}
	fp.Product = p
	images, err := ImagesByProductId(p.ProductId)
	if err != nil {
		return fp, fmt.Errorf("OneById(%v): %v", productId, err)
	}
	fp.Images = images
	return fp, nil
}

func validateQuery(queryParams []string, paramReg *regexp.Regexp) bool {
	if len(queryParams) == 0 {
		return false
	}
	for _, v := range queryParams {
		if !paramReg.MatchString(v) {
			return false
		}
	}
	return true
}

func getSortQuery(sortParams []string) string {
	isValid := validateQuery(sortParams, sortParamReg)
	if isValid {
		return strings.Join(sortParams, ", ")
	}
	return defaultOptions["sort"]
}

func getLimitQuery(limitParams []string) string {
	isValid := validateQuery(limitParams, limitParamReg)
	if isValid {
		return limitParams[0]
	}
	return defaultOptions["limit"]
}
