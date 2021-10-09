import React from "react";
import { ProductImage } from ".";

function Product({ product } = props) {
  return (
    <div className="catalog__item">
      <div className="cat-item">
        <div className="cat-item__img">{product.Images.map(image => <ProductImage key={image.ImageUrl} imageUrl={image.ImageUrl} />)}</div>
        <a href="#" className="cat-item__title">{product.Title}</a>
      </div>
    </div>
  );
}

export { Product }