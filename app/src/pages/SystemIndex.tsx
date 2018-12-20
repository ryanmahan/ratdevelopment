import * as React from "react";
import '../sass/custom-bulma.scss';
import { Divider } from "../components/layout/Divider";
import { PageTitle } from "../components/layout/PageTitle";
import { SystemIndexTable } from "../components/SystemIndexComponents/SystemIndexTable";
import { SearchBar } from "../components/SystemIndexComponents/SearchBar"

//import * as sample from "./SystemViewComponents/SampleSystem.json";
/* function updateState(snapshots) {
 *     this.setState({snapshots})
 * }
 *  */
export interface SystemIndexProps {

}
export interface SystemIndexState {
  snapshots: []
}
export class SystemIndex extends React.Component<SystemIndexProps, SystemIndexState> {

    constructor(props: SystemIndexProps){
        super(props);
        this.state = {
            snapshots: [],
        };
        this.handleSearchChange = this.handleSearchChange.bind(this)
    }

    componentDidMount() {
        this.getSnapshots("")
    }

    //fetch the latest snapshots and then update the state of the table.
    getSnapshots(searchstring: string) {
        //  Make the API call
        fetch(
            process.env.API_URL + "/api/tenants/1200944110/snapshots" + (searchstring ? "?searchString=" + searchstring : "")
        ).then(r => {
            //  When that returns convert it to json
            return r.json();
        }).then(j => {
            //  Finally set the state of the table to the list of snapshots returned
            this.setState({
                snapshots: j
            })
        });
    }

  handleSearchChange(query: string) {
      this.getSnapshots(encodeURI(query))
  }

  render() {
    return (
      <div className="container">
        <SearchBar handleSearchChange={this.handleSearchChange} />
        <PageTitle title={"Systems"} />
        <Divider />
        <SystemIndexTable snapshots={this.state.snapshots} />
      </div>
    )
  }
}
