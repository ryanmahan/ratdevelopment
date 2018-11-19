import * as React from "react";
import { Route, Switch } from "react-router-dom";
import { Navbar } from "./components/navigation/Navbar";
import { Divider } from "./components/layout/Divider";
import { Router } from "./components/navigation/AppRouter";

export const App = () => (
    <div>
        <Navbar fixed={false} active={0} navItems={[]}/>
        <Divider />
        <Switch>
            {Router.map((route, num) => (
                <Route key={num} path={route.route} exact={route.match} component={route.main} />
            ))}
        </Switch>
        <footer>&copy; 2018 HPE®</footer>
    </div>
);
