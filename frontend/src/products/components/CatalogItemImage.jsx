import React from "react";
import { catalogItemImageUrlType } from "./types";

CatalogItemImage.propTypes = {
  imageUrl: catalogItemImageUrlType.isRequired
}

function CatalogItemImage({ imageUrl }) {
  return <img src={imageUrl} alt="bla" />;
}

export { CatalogItemImage }