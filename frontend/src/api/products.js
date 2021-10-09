function fetchProducts() {
  return fetch("http://localhost:4000/react")
    .then(res => res.json())
}

export { fetchProducts };