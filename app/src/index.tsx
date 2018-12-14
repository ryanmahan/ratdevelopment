import * as React from "react";
import * as ReactDOM from "react-dom";
import { BrowserRouter } from "react-router-dom";
import { App } from "./App";

if (process.env.NODE_ENV == "development") {
    process.env.API_URL = "http://localhost:8081"
} else {
    process.env.API_URL = "http://35.231.27.158"
}
console.log(process.env.API_URL)
ReactDOM.render(
    <BrowserRouter>
        <App />
    </BrowserRouter>,
    document.getElementById("app-entry")
);
