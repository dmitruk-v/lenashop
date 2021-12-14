import PT from "prop-types";

const catalogItemImageUrlType = PT.string;

const catalogItemImageType = PT.shape({
  ImageUrl: PT.string
});

const catalogItemType = PT.shape({
  ProductId: PT.number,
  Title: PT.string,
  Quantity: PT.number,
  Price: PT.number,
  Images: PT.arrayOf(catalogItemImageType)
});

const catalogItemsType = PT.arrayOf(catalogItemType);

export {
  catalogItemImageUrlType, catalogItemImageType, catalogItemType, catalogItemsType
}