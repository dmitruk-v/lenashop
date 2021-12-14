import React from "react";
import PT from "prop-types";

CatalogOptions.propTypes = {
  onSort: PT.func,
  onLimit: PT.func
}

function CatalogOptions({ onSort, onLimit }) {
  const cookies = parseCookies(document.cookie);

  const sort = (evt) => {
    onSort("&sort=" + evt.target.value);
  }

  const limit = (evt) => {
    onLimit("&limit=" + evt.target.value);
  }

  return (
    <div className="catalog__options">
      <div className="catalog__opt">
        <label className="control-select">
          <select name="sort" onChange={sort} defaultValue={cookies["catalog-sort"]}>
            <option value="">Сортировка</option>
            <option value="price+asc">Цена: по возрастанию</option>
            <option value="price+desc">Цена: по убыванию</option>
          </select>
        </label>
      </div>
      <div className="catalog__opt">
        <label className="control-select">
          <select name="limit" onChange={limit} defaultValue={cookies["catalog-limit"]}>
            <option value="">Кол-во</option>
            <option value="4">4</option>
            <option value="8">8</option>
            <option value="12">12</option>
            <option value="16">16</option>
          </select>
        </label>
      </div>
    </div>
  );
}

/**
 * @function parseCookies
 * @param {string} cookiesStr
 * @return {Object.<string,string>}
 */
const parseCookies = (cookiesStr) => {
  const reg = /([A-Za-z0-9_-]+)=([^=;]*)/g;
  const result = {};
  let matches = reg.exec(cookiesStr);
  while (matches !== null) {
    result[matches[1]] = matches[2].replace(" ", "+");
    matches = reg.exec(cookiesStr);
  }
  return result;
}

export { CatalogOptions }