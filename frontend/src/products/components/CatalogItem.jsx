import React from "react";
import { CatalogItemImage } from ".";
import { catalogItemType } from "./types";

CatalogItem.propTypes = {
  product: catalogItemType.isRequired
}

function CatalogItem({ product }) {

  const renderStock = () => {
    if (product.Quantity > 10) {
      return <div className="cat-item__stock cat-item__stock--in">In stock</div>
    } else if (product.Quantity <= 10 && product.Quantity > 0) {
      return <div className="cat-item__stock cat-item__stock--ends">Out soon</div>
    } else {
      return <div className="cat-item__stock cat-item__stock--out">Out of stock</div>
    }
  }

  const titleHref = "/product/" + product.ProductId;

  return (
    <div className="cat-item">
      <a href={titleHref} className="cat-item__img">
        {product.Images.map(image => <CatalogItemImage key={image.ImageUrl} imageUrl={image.ImageUrl} />)}
      </a>
      <a href={titleHref} className="link cat-item__title">{product.Title}</a>
      <div className="cat-item__footer">
        <div className="cat-item__aside">
          <div className="cat-item__new-price">{product.Price} грн</div>
          {renderStock()}
        </div>
        {product.Quantity > 0 && (
          <div className="cat-item__buy">
            <form action="/cart/products/add" method="post">
              <input type="hidden" name="product_id" value={product.ProductId} />
              <input type="hidden" name="buy_quantity" value="1" />
              <button type="submit" className="button cat-item__btn-buy"></button>
            </form>
          </div>
        )}
      </div>
    </div>
  );
}

export { CatalogItem }