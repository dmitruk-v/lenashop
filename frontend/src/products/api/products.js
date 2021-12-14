const API_URL = "http://192.168.1.6:4000";

function fetchProducts(query, abortController) {
  /** @type {RequestInit} */
  const options = {
    headers: { "X-Requested-With": "XMLHttpRequest" },
    signal: abortController.signal,
    mode: "same-origin",
  };
  return fetch(API_URL + "/products?" + query, options)
    .then(res => res.json());
}

export { fetchProducts };