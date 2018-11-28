import * as React from "react";
import { SystemIndex } from "../../pages/SystemIndex";
import { Login } from "../../pages/Login/Login";
import {SystemView} from "../../pages/SystemView";

interface IRouter {
    route: string,
    match: boolean,
    main: any
}

export const Router: IRouter[] = [{
    route: "/",
    match: true,
    main: () => <SystemIndex />
}, {
    route: "/login",
    match: false,
    main: () => <Login />
}, {
    route: "/view/:serialNumber",
    match: false,
    main: SystemView
}];
