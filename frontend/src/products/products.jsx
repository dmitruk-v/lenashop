import React from "react";
import ReactDOM from "react-dom";
import { Catalog } from "./components/Catalog";

// we can try to create one redux instance for all site

ReactDOM.render(<Catalog />, document.querySelector(".catalog"));