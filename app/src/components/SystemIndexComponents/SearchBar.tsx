import * as React from "react";
import '../../sass/custom-bulma.scss';
import { Link } from "react-router-dom";
import { SystemIndexTable } from "./SystemIndexTable";

export interface SearchBarProps {
    handleSearchChange: any
}
export interface SearchBarState {
    searchString: string
}
export class SearchBar extends React.Component<SearchBarProps, SearchBarState> {
    constructor(props: SearchBarProps){
        super(props);
        this.state = {
            searchString: ""
        };
    }
    handleChange(event) {
        this.setState({ searchString: event.target.value})
    }
    handleSubmit() {
        this.props.handleSearchChange(this.state.searchString)
    }

  render() {
    return (
      <div className="SearchBar is-pulled-right">
        <form action="Search" method="get">
          <div className="field is-grouped">
            <div className="control">
              <input className="input is-normal" type="text" placeholder="Search.." value={this.state.searchString} onChange={this.handleChange} />
            </div>
            <div className="control">
                <button className="button is-primary" type="submit" onClick={this.handleSubmit}><i className="fa fa-search" ></i></button>
            </div>
          </div>
        </form>
      </div>
    )
  }
}
