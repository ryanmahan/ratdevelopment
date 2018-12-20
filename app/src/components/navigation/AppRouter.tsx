import * as React from "react";
import { Route, RouteProps, Redirect } from "react-router";
import { connect } from "react-redux";
import { SystemIndex } from "../../pages/SystemIndex";
import { Login } from "../../pages/Auth/Login/Login";
import {SystemView} from "../../pages/SystemView";
import { Error } from "../../pages/Misc/Error";
import { Logout } from "../../pages/Auth/Logout/Logout";
import { AppState } from "../../misc/state/reducers/Reducers";
import { About } from "../../pages/About/about";
import {
    AUTH_STATUS_LOGGEDIN,
    AUTH_STATUS_ANY,
    AUTH_STATUS_GUEST,
    AppAuthState
} from "../../misc/state/constants"; 

interface IRouter {
    route: string,
    visible: number,
    match: boolean,
    main: (...args: any[]) => JSX.Element
}

export const Router: IRouter[] = [{
    route: "/",
    visible: AUTH_STATUS_LOGGEDIN,
    match: true,
    main: (props: any) => <SystemIndex {...props} />
}, {
    route: "/login",
    visible: AUTH_STATUS_GUEST,
    match: false,
    main: () => <Login />
}, {
    route: "/logout",
    visible: AUTH_STATUS_LOGGEDIN,
    match: false,
    main: () => <Logout />
}, {
    route: "/view/:serialNumber",
    visible: AUTH_STATUS_LOGGEDIN,
    match: false,
    main: (props: any) => <SystemView {...props} />
}, {
    route: "/about",
    visible: AUTH_STATUS_ANY,
    match: false,
    main: () => <About />
}, {
    route: undefined,
    visible: AUTH_STATUS_ANY,
    match: undefined,
    main: () => <Error />
}];

interface CustomRouteProps extends RouteProps {
    key: number,
    visible: number,
    authState?: AppAuthState,
}

class CustomRouteComponent extends React.Component<CustomRouteProps, {}> {

    constructor(props: CustomRouteProps) {
        super(props);
    }

    render() {
        const { component: Component, authState, visible, ...args } = this.props;
        return (
            <Route {...args} render={(routeProps) =>
                ((authState.authenticated && visible === AUTH_STATUS_LOGGEDIN) ||
                    (!authState.authenticated && visible === AUTH_STATUS_GUEST) ||
                    (visible === AUTH_STATUS_ANY)) ? (
                        <Component {...routeProps} />
                    ) : ((authState.authenticated && visible === AUTH_STATUS_GUEST) ? (
                        <Redirect to={"/"} />) : (
                            <Redirect to={"/login"} />
                        )
                    )
                }
            />
        );
    }

}

const mapStateToProps = (state: AppState) => {
    return {
        authState: state.auth
    };
};

export const CustomRoute = connect(mapStateToProps)(CustomRouteComponent);
