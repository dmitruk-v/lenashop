import React from "react";
import PT from "prop-types";
import { CatalogItem } from "."
import { catalogItemsType } from "./types";

CatalogItems.propTypes = {
  products: catalogItemsType.isRequired,
  isLoadingProducts: PT.bool.isRequired,
};

function CatalogItems({ products, isLoadingProducts }) {
  return (
    <div className="catalog__items">
      {isLoadingProducts
        ? (
          <div className="loading-indicator">Loading products ...</div>
        ) : (
          <div className="catalog__grid">
            {products.map(p => {
              return <div className="catalog__item" key={p.ProductId}>
                <CatalogItem key={p.ProductId} product={p} />
              </div>
            })}
          </div>
        )
      }
    </div>
  );
}

export { CatalogItems }