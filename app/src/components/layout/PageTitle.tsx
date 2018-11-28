import * as React from "react";

export class PageTitle extends React.Component<{ title: string, extras: any[] }> {

    static defaultProps = {
        extras: {}
    };

    render() {

        let extras: any[] = [];
        for (let i: number = 0; i < this.props.extras.length; i++) {
            extras[i] = <div key={i} className="level-item">{this.props.extras[i]}</div>
        }
        return (
        <div className="level">
            <div className="level-left">
                <h2 className="title level-item">{this.props.title}</h2>
            </div>
            <div className="level-right">
                {extras}
            </div>
        </div>
        );
    }

}
