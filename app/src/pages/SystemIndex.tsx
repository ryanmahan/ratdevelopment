import * as React from "react";
import '../sass/custom-bulma.scss';
import {Divider} from "../components/layout/Divider";
import {PageTitle} from "../components/layout/PageTitle";
import {SystemIndexTable} from "../components/SystemIndexComponents/SystemIndexTable";
import {AppAuthState} from "../misc/state/constants";
import {SearchBar} from "../components/SystemIndexComponents/SearchBar"
import {authorizedFetch} from "../misc/state/constants";
import * as queryString from "query-string";
import * as moment from 'moment';
import { RouteComponentProps } from "react-router";

//import * as sample from "./SystemViewComponents/SampleSystem.json";
/* function updateState(snapshots) {
 *     this.setState({snapshots})
 * }
 *  */
export interface SystemIndexProps {
    authState: AppAuthState
}
export interface SystemIndexState {
  snapshots: any[]
}
export class SystemIndex extends React.Component<SystemIndexProps & RouteComponentProps, SystemIndexState> {

    constructor(props: SystemIndexProps & RouteComponentProps){
        super(props);
        this.state = {
            snapshots: [],
        };
        this.handleSearchChange = this.handleSearchChange.bind(this)
        this.handleFilterChange = this.handleFilterChange.bind(this)
    }

    componentDidMount() {
        this.getSnapshots("")
    }

    //fetch the latest snapshots and then update the state of the table.
    getSnapshots(searchstring: string) {
        //  Make the API call
        authorizedFetch(
            "/api/tenants/hpe/snapshots" + (searchstring ? "?searchString=" + searchstring : ""),
            this.props.authState
        ).then(r => {
            //  When that returns convert it to json
            return r.json();
        }).then(j => {
            //  Finally set the state of the table to the list of snapshots returned
            j = j.map(o => {if(!o.system["companyName"]) {o.system["companyName"] = ""} return o})
            this.setState({
                snapshots: j
            }, () => {
                const values = queryString.parse(this.props.location.search);
                this.performSort(values);
            });
        });
    }

  handleSearchChange(query: string) {
      this.getSnapshots(encodeURI(query))
  }


    componentDidUpdate(prevProps: any) { // fires all subsequent times except first
        const values = queryString.parse(this.props.location.search);
        const prevValues = queryString.parse(prevProps.location.search);
        if ((values.filter !== prevValues.filter) || (values.order !== prevValues.order)) {
            // determines if anything has been updated
            this.performSort(values);
        }
    }

    performSort(values: any) {
        const filterColumn = values.filter;
        const filterOrder = values.order;

        if (filterColumn != null && filterOrder != null) {
            let sortOrder: number = 0;
            if (filterColumn === "serial") {
                if (filterOrder === "asc") {
                    // if ascending
                    sortOrder = 0;
                } else if (filterOrder === "desc") {
                    // if descending
                    sortOrder = 1;
                }
                this.onSortSerial("serialNumberInserv", sortOrder);
            } else if (filterColumn === "company") {
                if (filterOrder === "asc") {
                    sortOrder = 0;
                } else if (filterOrder === "desc") {
                    sortOrder = 1;
                }
                this.onSortCompany("companyName", sortOrder);
            } else if (filterColumn === "capacity") {
                if (filterOrder === "asc") {
                    sortOrder = 0;
                } else if (filterOrder === "desc") {
                    sortOrder = 1;
                }
                this.onSortCapacity("freeTiB", "sizeTiB", sortOrder);
            } else if (filterColumn === "date") {
                if (filterOrder === "asc") {
                    sortOrder = 0;
                } else if (filterOrder === "desc") {
                    sortOrder = 1;
                }
                this.onSortDate("date", sortOrder);
            }
        }
    }

    handleFilterChange(columnName: string) {
        const values = queryString.parse(this.props.location.search);
        const filterOrder = values.order;
        if (filterOrder === "asc") { // if current order is ascending
            this.props.history.push("/?filter=" + columnName + "&order=desc"); // clicking sort with make order descending
        }
        else {
            this.props.history.push("/?filter=" + columnName + "&order=asc");
        }
    }

    //onSort methods take the button click event and will either a)sort the list if unsorted or b)reverse the list if already sorted
    onSortSerial(sortKey: any, order: number) {
        const snapshots = this.state.snapshots;
        if (order === 0) {
            snapshots.sort((a, b) => a[sortKey].localeCompare(b[sortKey]));
        } else {
            snapshots.sort((a, b) => b[sortKey].localeCompare(a[sortKey]));
        }
        this.setState({ snapshots });
    }

    onSortCompany(sortKey: any, order: number) {
        const snapshots = this.state.snapshots;
        if (order === 0){
            snapshots.sort((a, b) => a.system[sortKey].localeCompare(b.system[sortKey]));
        } else {
            snapshots.sort((a, b) => b.system[sortKey].localeCompare(a.system[sortKey]));
        }
        this.setState({ snapshots });
    }

    onSortCapacity(sortKey: any, sortKey2: any, order: number) {
        const snapshots = this.state.snapshots;
        if (order === 0) {
            snapshots.sort((a, b) => ((b.capacity.total[sortKey]/b.capacity.total[sortKey2]).toString().localeCompare((a.capacity.total[sortKey]/a.capacity.total[sortKey2]).toString())));
        } else {
            snapshots.sort((a, b) => ((a.capacity.total[sortKey]/a.capacity.total[sortKey2]).toString().localeCompare((b.capacity.total[sortKey]/b.capacity.total[sortKey2]).toString())));
        }
        this.setState({ snapshots });
    }

    onSortDate(sortKey: any, order: number) {
        const snapshots = this.state.snapshots;
        if (order === 0) {
            snapshots.sort((a, b) => a[sortKey].localeCompare(b[sortKey]));
        } else {
            snapshots.sort((a, b) => b[sortKey].localeCompare(a[sortKey]));
        }
        this.setState({ snapshots });
    }


  render() {
    return (
      <div className="container">
        <SearchBar handleSearchChange={this.handleSearchChange}  />
        <PageTitle title={"Systems"} />
        <Divider />
        <SystemIndexTable authState={this.props.authState} snapshots={this.state.snapshots} handleFilterChange={this.handleFilterChange}/>
      </div>
    )
  }
}
