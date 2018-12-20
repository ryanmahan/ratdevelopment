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
        this.handleChange = this.handleChange.bind(this)
        this.handleSubmit = this.handleSubmit.bind(this)
    }
    handleChange(event) {
        this.setState({ searchString: event.target.value})
    }
    handleSubmit(event) {
        event.preventDefault()
        this.props.handleSearchChange(this.state.searchString)
    }

  render() {
    return (
      <div className="SearchBar is-pulled-right">
        <form action="Search" method="get" onSubmit={this.handleSubmit.bind(this)}>
          <div className="field is-grouped">
            <div className="control">
              <input className="input is-normal" type="text" placeholder="Search.." value={this.state.searchString} onChange={this.handleChange} />
            </div>
            <div className="control">
                <button className="button is-primary" type="submit" ><i className="fa fa-search" ></i></button>
            </div>
          </div>
        </form>
      </div>
    )
  }
}
