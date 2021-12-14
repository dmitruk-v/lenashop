import React from "react";
import { useProducts } from "../hooks";
import { CatalogItems } from "./CatalogItems";
import { CatalogOptions } from "./CatalogOptions";

function Catalog() {
  const { products, load, isLoadingProducts } = useProducts({
    initialLoad: true,
    // initialQuery: "sort=price+asc&limit=12"
  });

  const handleSort = (sortQuery) => {
    load(sortQuery);
  }

  const handleLimit = (limitQuery) => {
    load(limitQuery);
  }

  return (
    <>
      <CatalogOptions onSort={handleSort} onLimit={handleLimit} />
      <CatalogItems products={products} isLoadingProducts={isLoadingProducts} />
      <div className="catalog__more">mooooooore items...</div>
    </>
  )
}

export { Catalog }