import * as React from "react";
import '../../sass/custom-bulma.scss';
import {AppAuthState} from "../../misc/state/constants";
import {Link, withRouter, RouteComponentProps} from "react-router-dom";
import {API_URL} from "../../misc/state/constants";
import * as queryString from "query-string";
import * as moment from 'moment';

//reduced getters to only the helpful ones
function getCapacity(currentRow: any) {
  return 100 - Math.trunc(100 * (currentRow.capacity.total.freeTiB / currentRow.capacity.total.sizeTiB));
}

function getWarningImage(currentRow: any) {
  if (Math.trunc(100 * (currentRow.capacity.total.freeTiB / currentRow.capacity.total.sizeTiB)) <= 30) {
    return <figure className="image is-24x24 is-pulled-right">
      <img src="https://img.icons8.com/color/50/000000/high-priority.png" alt="Warning Low Capacity" title="Warning: Capacity below 30%"></img>
    </figure>;
  }
}

export interface SystemIndexTableProps {
  snapshots: any[]
  authState: AppAuthState
  history?: { push(path: string): void }
  handleFilterChange: any
}

interface SystemIndexTableState {
}

//exports the actual table component which calls our fillArray method
class SystemIndexTableComponent extends React.Component<SystemIndexTableProps & RouteComponentProps, SystemIndexTableState> {


    //the table is rendered
    //instead of an outside function iterating, the body is now rendered in a callback function using map
    //buttons will call sort functions
    render() {
        let newData = this.props.snapshots;
        if ( !newData ) return null
        return (
            <table id={"myTable"} className="table is-fullwidth is-bordered is-striped">
                <thead>
                <tr>
                    <th>Serial Number
                        <button className ="button is-dark is-small is-pulled-right" onClick={e => this.props.handleFilterChange("serial")}>
                            <i className="fas fa-sort"></i>
                        </button>
                    </th>
                    <th>Company
                        <button className ="button is-dark is-small is-pulled-right" onClick={e => this.props.handleFilterChange("company")}>
                            <i className="fas fa-sort"></i>
                        </button>
                    </th>
                    <th>Data Used
                        <button className ="button is-dark is-small is-pulled-right" onClick={e => this.props.handleFilterChange("capacity")}>
                            <i className="fas fa-sort"></i>
                        </button>
                    </th>
                    <th >Last Updated
                        <button className ="button is-dark is-small is-pulled-right" onClick={e => this.props.handleFilterChange("date")}>
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


}

export const SystemIndexTable = withRouter(SystemIndexTableComponent)
