import * as React from "react";
import {Link, NavLink} from "react-router-dom";
import '../../sass/custom-bulma.scss';

export interface NavbarProps {
    navItems: string[],
    active: number,
    fixed: boolean
}

export class Navbar extends React.Component<NavbarProps> {

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

        return (
            <nav className={"navbar" + fixed} role="navigation" aria-label="main navigation">
                <div className="container">
                    <div className="navbar-brand">
                        <Link className="navbar-item" to="/">

                                <img src="https://upload.wikimedia.org/wikipedia/commons/4/46/Hewlett_Packard_Enterprise_logo.svg"/>

                        </Link>
                    </div>
                    <div className="navbar-end">
                        <Link className="navbar-item level" to="/login">
                            <span style={{paddingRight: "0.5rem"}}>Account</span>
                            <span className="icon">
                                <i className="fas fa-user-circle fa-2x"/>
                            </span>
                        </Link>
                    </div>
                </div>
            </nav>
        );
    }
}
