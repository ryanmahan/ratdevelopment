import * as React from "react";
import '../../sass/custom-bulma.scss';
import { Link } from "react-router-dom";
import { SystemIndexTable } from "./SystemIndexTable";

export class SearchBar extends React.Component {


	render() {
		return (
			<div className="SearchBar is-pulled-right">
				<form action="Search" method="get">
					<div className="field is-grouped">
						<div className="control">
							<input className="input is-normal" type="text" placeholder="Search.." />
						</div>
						<div className="control">
							<button className="button is-primary" type="submit"><i className="fa fa-search"></i></button>
						</div>
					</div>
				</form>
			</div>
		)
	}
}