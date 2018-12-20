import * as React from "react";
import { connect } from "react-redux";
import { withRouter } from "react-router-dom";
import { AppAuthState } from "../../../misc/state/constants";
import { setState } from "../../../misc/state/actions/Actions";
import { API_URL } from "../../../misc/state/constants"
import "./Login.css";
import auth0 from 'auth0-js';

interface loginProps {
    loginUser(args: AppAuthState): void,
    history?: { push(path: string): void }
}

interface loginState {
    userName: string,
    password: string
}

class LoginComponent extends React.Component<loginProps, loginState> {
    constructor(props: any) {
        super(props);
        this.state = { userName: "", password: ""};
        this.onLogin = this.onLogin.bind(this);
    }
    auth = new auth0.WebAuth({
      domain: 'rat-dev.auth0.com',
      clientID: 'kGXtSueZuisYoZneXoOOZUm_jJs33lhp',
      responseType: 'token id_token',
      redirectUri: API_URL + '/login',
      scope: 'openid'
    });

    componentDidMount(){
        document.title = "Sign in | HPE® Official Site";
        this.handleAuthentication()
    }

    login(usernameVar, passwordVar) {
      this.auth.login({
          realm: 'Username-Password-Authentication',
          email: usernameVar,
          password: passwordVar
        }, function(err) {
        if (err) {
          console.log(err['description']);
          alert('Error: ' + err['description']);
        }
      });
    }

    handleAuthentication() {
      this.auth.parseHash((err, authResult) => {
        if (authResult && authResult.accessToken && authResult.idToken) {
          this.setSession(authResult);
        } else if (err) {
          console.log(err);
          alert(`Error: ${err.error}. Check the console for further details.`);
        }
      });
    }

    setSession(authResult) {
      // Set the time that the access token will expire at
      this.props.loginUser({
          authenticated: true,
          access_token: authResult.accessToken,
          userId: authResult.idToken,
          userName: localStorage.getItem('email')
      });
      this.props.history.push("/");
      // navigate to the home route
    }

    onLogin(event: React.FormEvent<HTMLFormElement>) {
        event.preventDefault();
        const {userName, password} = this.state;
        if (!(userName && password)) {
          return;
        }
        localStorage.setItem('email', userName);
        this.login(userName, password)

    }

    render() {
        return <div className="log-in">
            <div className="container">
                <form className="login-form" onSubmit={e => this.onLogin(e)}>
                    <h1 className="log-in-two">Sign in</h1>
                    <span className="required">Required *</span>
                    <div className="free-line"/>
                    <div>
                        <label>User ID </label>
                        <span className="required">*</span>
                    </div>
                    <div>
                        <input className="textbox" required onChange={e => this.setState({ userName: e.target.value})} />
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
                        <input type="password" className="textbox" required onChange={e => this.setState({ password: e.target.value})} />
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
                        <input className="button-style signin-style" type="submit" value="Sign in" />
                    </div>
                </form>
            </div>
        </div>
    }
}

const mapDispatchToProps = (dispatcher: any) => {
    return {
        loginUser: (args: AppAuthState) => {
            dispatcher(setState(args));
        }
    }
};

export const Login = connect(null, mapDispatchToProps)(withRouter(LoginComponent));
