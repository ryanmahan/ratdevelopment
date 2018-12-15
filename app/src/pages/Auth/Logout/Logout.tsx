import * as React from "react";
import { withRouter } from "react-router-dom";
import { UserService } from "../../../misc/state/service/User";

interface logoutProps {
    history: { push(path: string): void }
}

class LogoutComponent extends React.Component<logoutProps, {}> {
    constructor(props: any) {
        super(props);
    }

    componentWillMount() {
        UserService.logout();
        window.location.href = "/";
    }

    render() {
        return "";
    }
}

export const Logout = withRouter(LogoutComponent);
