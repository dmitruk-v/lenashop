import { useEffect, useState } from "react";
import { fetchProducts } from "../api/products";

/**
 * @typedef {Object} Options
 * @property {boolean} initialLoad
 * @property {string} initialQuery
 * 
 * @function useProducts
 * @param {Options} options
 */
function useProducts(options = { initialLoad: false, initialQuery: "" }) {
  const [products, setProducts] = useState([]);
  const [isLoading, setIsLoading] = useState(false);
  const abortController = new AbortController();

  useEffect(() => {
    if (options.initialLoad) {
      load(options.initialQuery);
    }
  }, []);

  const load = (query = "") => {
    setIsLoading(true);
    fetchProducts(query, abortController)
      .then(products => {
        setProducts(products);
        setIsLoading(false);
      })
      .catch(err => {
        console.error("[useProducts]:", err)
      });
  }

  return {
    products: products,
    load: load,
    isLoadingProducts: isLoading,
    cancel: abortController.abort,
  };
}

export { useProducts };