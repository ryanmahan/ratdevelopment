import * as React from "react";
import '../../sass/custom-bulma.scss';
import {Link, withRouter, RouteComponentProps} from "react-router-dom";
import {API_URL} from "../../misc/state/constants"
import * as queryString from "query-string";
import * as moment from 'moment';

//reduced getters to only the helpful ones
function getCapacity(currentRow: any){
    return 100 - Math.trunc(100 * (currentRow.capacity.total.freeTiB / currentRow.capacity.total.sizeTiB));
}

function getWarningImage(currentRow: any) {
    if(Math.trunc(100 * (currentRow.capacity.total.freeTiB / currentRow.capacity.total.sizeTiB)) <= 30){
        return <figure className="image is-24x24 is-pulled-right">
            <img src="https://img.icons8.com/color/50/000000/high-priority.png" alt="Warning Low Capacity" title= "Warning: Capacity below 30%"></img>
        </figure>;
    }
}

export interface SystemIndexTableProps {
    history?: { push(path: string): void }
}

interface SystemIndexTableState {
    snapshots: any[]
}

//exports the actual table component which calls our fillArray method
class SystemIndexTableComponent extends React.Component<SystemIndexTableProps & RouteComponentProps, SystemIndexTableState> {

    constructor(props: SystemIndexTableProps & RouteComponentProps) {
        super(props);
        this.state = { snapshots: [] };
        this.onSortSerial = this.onSortSerial.bind(this);
        this.onSortCompany = this.onSortCompany.bind(this);
        this.onSortCapacity = this.onSortCapacity.bind(this);
        this.onSortDate = this.onSortDate.bind(this);
    }

    // This is triggered when this component is mounted
    componentDidMount() { // fires only the first time
        this.getSnapshots(() => {
            const values = queryString.parse(this.props.location.search);
            this.performSort(values);
        })
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

    applyFilter(columnName: string) {
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

    //the table is rendered
    //instead of an outside function iterating, the body is now rendered in a callback function using map
    //buttons will call sort functions
    render() {
        let newData = this.state.snapshots;
        return (
            <table id={"myTable"} className="table is-fullwidth is-bordered is-striped">
                <thead>
                <tr>
                    <th>Serial Number
                        <button className ="button is-dark is-small is-pulled-right" onClick={e => this.applyFilter("serial")}>
                            <i className="fas fa-sort"></i>
                        </button>
                    </th>
                    <th>Company
                        <button className ="button is-dark is-small is-pulled-right" onClick={e => this.applyFilter("company")}>
                            <i className="fas fa-sort"></i>
                        </button>
                    </th>
                    <th>Data Used
                        <button className ="button is-dark is-small is-pulled-right" onClick={e => this.applyFilter("capacity")}>
                            <i className="fas fa-sort"></i>
                        </button>
                    </th>
                    <th >Last Updated
                        <button className ="button is-dark is-small is-pulled-right" onClick={e => this.applyFilter("date")}>
                            <i className="fas fa-sort"></i>
                        </button>
                    </th>
                </tr>
                </thead>
                <tbody>
                {
                    newData.map(function (system, index) {
                        return(
                            <tr key={index} data-item="Serial Number">

                                <td><Link to={"/view/" + system.serialNumberInserv} >
                                    {system.serialNumberInserv}
                                    </Link></td>
                                <td>{system.system.companyName}</td>
                                <td>
                                    <div className="columns">
                                        <div className="column">
                                            <progress className="progress is-primary" value={getCapacity(system)} max="100">
                                            {getCapacity(system)}
                                            </progress>
                                        {getCapacity(system)}% used
                                        </div>
                                        <div className="column is-2">
                                            {getWarningImage(system)}
                                        </div>
                                    </div>
                                </td>
                                <td>{moment(system.date).utc().format('MMMM Do YYYY, h:mm A')}</td>
                            </tr>
                        );
                    })
                }
                </tbody>
            </table>
        )
    }

    //fetch the latest snapshots and then update the state of the table.
    getSnapshots(cb: any) {
        //  Make the API call
        fetch(
            API_URL + "/api/tenants/1200944110/snapshots"
        ).then(r => {
            //  When that returns convert it to json
            return r.json();
        }).then(j => {
            //  Finally set the state of the table to the list of snapshots returned
            this.setState({
                snapshots: j
            });
            cb();
        });
    }
}

export const SystemIndexTable = withRouter(SystemIndexTableComponent)

