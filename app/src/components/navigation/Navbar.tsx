import * as React from "react";
import { Link } from "react-router-dom";
import '../../sass/custom-bulma.scss';
import { AppAuthState } from "../../misc/state/constants";
import { connect } from "react-redux";
import { AppState } from "../../misc/state/reducers/Reducers";

export interface NavbarProps {
    authState?: AppAuthState,
    navItems: string[],
    active: number,
    fixed: boolean
}

interface INavbarState {
    toggled: boolean
}

class NavbarComponent extends React.Component<NavbarProps, INavbarState> {

    constructor(props: any) {
        super(props);
        this.setDropdown = this.setDropdown.bind(this);
        this.state = {
            toggled: false
        };
    }

    setDropdown(state: boolean) {
        console.log("toggle drop");
        this.setState({
            toggled: state
        });
    }

    render() {
        let items: any[] = [];
        for (let i: number = 0; i < this.props.navItems.length; i++) {
            let active: string = "";
            if (i == this.props.active) {
                active = " is-active";
            }
            items[i] = <a key={i} className={"navbar-item" + active}>
                {this.props.navItems[i]}
            </a>;
        }

        let fixed: string = "";
        if (this.props.fixed == true) {
            fixed = " is-fixed-top";
        }

        let displayUser = "Account";
        if (this.props.authState.authenticated) {
            displayUser = this.props.authState.userName;
        }

        return (
            <nav className={"navbar" + fixed} role="navigation" aria-label="main navigation">
                <div className="container">
                    <div className="navbar-brand">
                        <Link className="navbar-item" to="/">
                            <img src="https://upload.wikimedia.org/wikipedia/commons/4/46/Hewlett_Packard_Enterprise_logo.svg"/>
                        </Link>
                        <Link className="navbar-item" to="/about">About</Link>
                    </div>
                    <div className="navbar-end">
                        <div className={"navbar-item level dropdown " + (this.state.toggled ? "is-active" : "")}>
                            <div className="dropdown-trigger" onClick={() => this.setDropdown(!this.state.toggled)}>
                                <a className="navbar-item level" aria-haspopup="true" aria-controls="dropdown-menu">
                                    <span style={{paddingRight: "0.5rem"}}>{displayUser}</span>
                                    <span className="icon">
                                        <i className="fas fa-user-circle fa-2x"/>
                                    </span>
                                </a>
                            </div>
                            <div className="dropdown-menu" id="dropdown-menu" role="menu">
                                {this.props.authState.authenticated &&
                                    <div className="dropdown-content">
                                        <Link to="/logout" className="dropdown-item">Logout</Link>
                                    </div>
                                }
                                {!this.props.authState.authenticated &&
                                    <div className="dropdown-content">
                                        <Link to="/login" className="dropdown-item">Login</Link>
                                    </div>
                                }
                            </div>
                        </div>
                    </div>
                </div>
            </nav>
        );
    }
}

const mapStateToProps = (state: AppState) => {
    return {
        authState: state.auth
    };
};

export const Navbar = connect(mapStateToProps)(NavbarComponent);
