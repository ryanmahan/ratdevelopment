import * as React from "react";
import '../sass/custom-bulma.scss';
import { Divider } from "../components/layout/Divider";
import { PageTitle } from "../components/layout/PageTitle";
import { SystemIndexTable } from "../components/SystemIndexComponents/SystemIndexTable";
import { SearchBar } from "../components/SystemIndexComponents/SearchBar"

//import * as sample from "./SystemViewComponents/SampleSystem.json";

export interface SystemIndexProps {

}
export interface SystemIndexState {
  snapshotArray: any[]
  searchString: string
}
export class SystemIndex extends React.Component<SystemIndexProps, SystemIndexState> {
  handleSearchChange(query: string) {
      this.setState({ snapshotArray: [], searchString: query });
      fetch(
          process.env.API_URL + "/api/tenants/1200944110/snapshots" + this.state.searchString ? "?searchString=" + this.state.searchString : ""
      ).then(r => {
          //  When that returns convert it to json
          return r.json();
      }).then(j => {
          //  Finally set the state of the table to the list of snapshots returned
          this.setState({
              snapshotArray: j,
              searchString: this.state.searchString
          })
      });
  }

  render() {
    return (
      <div className="container">
        <SearchBar handleSearchChange={this.handleSearchChange} />
        <PageTitle title={"Systems"} />
        <Divider />
        <SystemIndexTable searchstring={this.state.searchString} />
      </div>
    )
  }
}
