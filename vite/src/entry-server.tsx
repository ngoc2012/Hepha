import React from "react";
import ReactDOMServer from "react-dom/server";
import { StaticRouter } from "react-router-dom/server";
import SSR from "./SSR";
import "./index.css";

export function render(url: string) {
  return ReactDOMServer.renderToString(
    <React.StrictMode>
      <StaticRouter location={url}>
        <SSR />
      </StaticRouter>
    </React.StrictMode>,
  );
}
