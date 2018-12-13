import * as React from "react";
import { Switch } from "react-router-dom";
import { Navbar } from "./components/navigation/Navbar";
import { Divider } from "./components/layout/Divider";
import { Router, CustomRoute } from "./components/navigation/AppRouter";

export const App = () => (
    <div>
        <Navbar fixed={false} active={0} navItems={[]}/>
        <Divider />
        <Switch>
            {Router.map((route, num) => (
                <CustomRoute key={num} path={route.route} exact={route.match} component={route.main} visible={route.visible} />
            ))}
        </Switch>
        <footer>&copy; 2018 HPE®</footer>
    </div>
);
