import React from "react";

function ProductImage({ imageUrl } = props) {
  return <img src={imageUrl} alt="bla" />;
}

export { ProductImage }