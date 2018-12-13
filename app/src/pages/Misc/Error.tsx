import * as React from "react";
import "../../sass/custom-bulma.scss";
import "./Error.css";

export class Error extends React.Component<{}, {}> {
    constructor(props: any) {
        super(props);
    }

    render() {
        return <div className="container">
            <h1 className="head">Page Not Found</h1>
            <p className="describe-error">Nothing to see here...keep moving.</p>
        </div>
    }
}