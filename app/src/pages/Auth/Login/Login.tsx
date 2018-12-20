import * as React from "react";
import { connect } from "react-redux";
import { withRouter } from "react-router-dom";
import { AppAuthState } from "../../../misc/state/constants";
import { setState } from "../../../misc/state/actions/Actions";
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
        scope: 'openid profile read:all',
        audience: 'https://mousefb/api',
        redirectUri: 'http://localhost:8080/login'
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

    signup(usernameVar, passwordVar) {
      this.auth.signup({
          connection: 'Username-Password-Authentication',
          email: usernameVar,
          password: passwordVar
        }, function(err) {
        if (err) {
          console.log(err['description']);
          alert('Error: ' + err['description']);
        }
        return alert('Successfully Registered!')
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

    onLogin(event) {
        event.preventDefault();
        const {userName, password} = this.state;
        if (!(userName && password)) {
          return;
        }

        localStorage.setItem('email', userName);
        this.login(userName, password)

    }
    onSignup(event) {
        event.preventDefault();
        const {userName, password} = this.state;
        if (!(userName && password)) {
          return;
        }
        console.log("signed up")
        localStorage.setItem('email', userName);
        this.signup(userName, password)

    }

    render() {
        return <div className="log-in">
            <div className="container">
                <form className="login-form">
                    <h1 className="log-in-two">Login</h1>
                    <span className="required">Required *</span>
                    <div className="free-line"/>
                    <div>
                        <label>Email</label>
                        <span className="required">*</span>
                    </div>
                    <div>
                        <input className="textbox" required onChange={e => this.setState({ userName: e.target.value})} />
                    </div>
                    <br/>
                    <div>
                        <label>Password </label>
                        <span className="required">*</span>
                    </div>
                    <div>
                        <input type="password" className="textbox" required onChange={e => this.setState({ password: e.target.value})} />
                    </div>

                    <br/>
                    <div className="login-form-action">
                        <input className="button-style" type="submit" value="Create Account"  onClick={f => this.onSignup(f)}/>
                        <input className="button-style signin-style" type="submit" value="Login"  onClick={e => this.onLogin(e)}/>
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
