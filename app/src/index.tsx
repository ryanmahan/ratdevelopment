import * as React from "react";
import * as ReactDOM from "react-dom";
import { BrowserRouter } from "react-router-dom";
import { createStore } from "redux";
import { Provider } from "react-redux";
import { App } from "./App";
import { AppReducer } from "./misc/state/reducers/Reducers";

let state: Object = {};
if (localStorage.getItem("state")) {
    state = JSON.parse(localStorage.getItem("state"));
}

const appStore = createStore(AppReducer, state);
appStore.subscribe(() => {
    localStorage.setItem("state", JSON.stringify(appStore.getState()));
});

if (process.env.NODE_ENV == "development") {
    process.env.API_URL = "http://localhost:8081"
} else {
    process.env.API_URL = "http://35.231.27.158"
}
console.log(process.env.API_URL)
ReactDOM.render(
    <Provider store={appStore}>
        <BrowserRouter>
            <App />
        </BrowserRouter>
    </Provider>,
    document.getElementById("app-entry")
);
