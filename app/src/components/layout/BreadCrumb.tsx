import * as React from "react";

export interface BreadCrumbProps {
    labels: string[],
    links: string[],
    final: string
}

export class BreadCrumb extends React.Component<BreadCrumbProps> {
    static defaultProps = {
        links: {}
    };

    render() {
        let crumbs: any[] = [];
        for (let i: number = 0; i < this.props.labels.length; i++) {
            let link = this.props.links.length > i ? this.props.links[i] : "#";
            crumbs.push(
                <li key={i}><a href={link}>{this.props.labels[i]}</a></li>
            )
        }
        return (
            <nav className="breadcrumb">
                <ul>
                    {crumbs}
                    <li className="is-active"><a href="#" aria-current="page">{this.props.final}</a></li>
                </ul>
            </nav>
        );
    }
}