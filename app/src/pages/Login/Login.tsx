import * as React from "react";
import "./Login.css";


export class Login extends React.Component<{},{}> {
    constructor(props: any) {
        super(props);
    }

    componentDidMount(){
        document.title = "Sign in | HPE® Official Site"
    }

    //TODO: routing system needs to be set up in order to be able to use Link for "forgot user ID" and "forgot password"
    render() {
        return <div className="log-in">
            <div className="container">
                <div className="login-form">
                    <h1 className="log-in-two">Sign in</h1>
                    <span className="required">Required *</span>
                    <div className="free-line"/>
                    <div>
                        <label>User ID </label>
                        <span className="required">*</span> 
                    </div>
                    <div>
                        <input className="textbox" required/>
                    </div>
                    <div className="user-ID">
                        <span className="user-ID-reminder">Your user ID may be your email.</span>
                        <a className="user-ID-forgot-ID" href="#">
                            <span>Forgot User ID</span>
                        </a>
                    </div>
                    <div>
                        <label>Password </label>
                        <span className="required">*</span> 
                    </div>
                    <div>
                        <input type="password" className="textbox" required/>
                    </div>
                    <div className = "user-ID-forgot-pswd">
                        <a className= "user-ID-forgot-pswd-text" href="#">
                            <span>Forgot Password</span>
                        </a>
                    </div>
                    <div className="checkbox">
                        <label>
                            <input className="remember-me" type="checkbox"/>Remember me on this computer
                        </label>
                    </div>
                    <div className="login-form-action">
                        <button className="button-style" type="button">Create an account</button>
                        <button className="button-style signin-style" type="button">Sign in</button>
                    </div>
                </div>
            </div>
        </div>
    }
}