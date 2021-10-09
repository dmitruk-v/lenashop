import { useEffect, useState } from "react";
import { fetchProducts } from "../../api/products";

function useProducts() {
  const [products, setProducts] = useState([]);
  const [isLoading, setIsLoading] = useState(false);

  const handleFetch = () => {
    setIsLoading(true);
    fetchProducts().then(products => {
      setProducts(products);
      setIsLoading(false);
    });
  }

  useEffect(() => {
    handleFetch();
  }, []);

  return {
    products: products,
    refresh: handleFetch,
    isLoadingProducts: isLoading,
  };
}

export { useProducts };