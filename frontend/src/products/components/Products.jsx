import React from "react";
import { useProducts } from "../hooks";
import { Product } from "."

function Products(props) {
  const { products, refresh, isLoadingProducts } = useProducts();

  return (<>
    <div className="catalog__grid">
      {
        isLoadingProducts
          ? <div className="loading-indicator">Loading products ...</div>
          : products.map(p => <Product key={p.ProductId} product={p} />)
      }
    </div>
    <button onClick={() => refresh()}>REFRESH</button>
  </>);
}

export { Products }