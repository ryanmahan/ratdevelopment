import * as React from "react";
import '../../sass/custom-bulma.scss';
import { Link } from "react-router-dom";

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
  searchstring: string
}

interface SystemIndexTableState {
  snapshots: any[]
}

//exports the actual table component which calls our fillArray method
export class SystemIndexTable extends React.Component<SystemIndexTableProps, SystemIndexTableState> {

  constructor(props: SystemIndexTableProps) {
    super(props);
    this.state = { snapshots: [] };
    this.onSortSerial = this.onSortSerial.bind(this);
    this.onSortCompany = this.onSortCompany.bind(this);
    this.onSortCapacity = this.onSortCapacity.bind(this);
    this.onSortDate = this.onSortDate.bind(this);
  }

  //This is triggered when this component is mounted
  componentDidMount() {
    this.getSnapshots(this.props.searchstring)
  }

  //variables which keep track of the states ordering
  serialOrder = 0;
  companyOrder = 0;
  capacityOrder = 0;
  dateOrder = 0;

  //onSort methods take the button click event and will either a)sort the list if unsorted or b)reverse the list if already sorted
  onSortSerial(event: any, sortKey: any) {
    const snapshots = this.state.snapshots;
    if (this.serialOrder == 0) {
      snapshots.sort((a, b) => a[sortKey].localeCompare(b[sortKey]));
      this.serialOrder = 1;
      this.companyOrder = 0;
      this.capacityOrder = 0;
      this.dateOrder = 0;
    }
    else {
      snapshots.reverse();
    }
    this.setState({ snapshots });
  }

  onSortCompany(event: any, sortKey: any) {
    const snapshots = this.state.snapshots;
    if (this.companyOrder == 0) {
      snapshots.sort((a, b) => a.system[sortKey].localeCompare(b.system[sortKey]));
      this.serialOrder = 0;
      this.companyOrder = 1;
      this.capacityOrder = 0;
      this.dateOrder = 0;
    }
    else {
      snapshots.reverse();
    }
    this.setState({ snapshots });
  }

  onSortCapacity(event: any, sortKey: any, sortKey2: any) {
    const snapshots = this.state.snapshots;
    if (this.capacityOrder == 0) {
      snapshots.sort((a, b) => ((a.capacity.total[sortKey] / a.capacity.total[sortKey2]).toString().localeCompare((b.capacity.total[sortKey] / b.capacity.total[sortKey2]).toString())));
      this.serialOrder = 0;
      this.companyOrder = 0;
      this.capacityOrder = 1;
      this.dateOrder = 0;
    }
    else {
      snapshots.reverse();
    }
    this.setState({ snapshots });
  }

  onSortDate(event: any, sortKey: any) {
    const snapshots = this.state.snapshots;
    if (this.dateOrder == 0) {
      snapshots.sort((a, b) => a[sortKey].localeCompare(b[sortKey]));
      this.serialOrder = 0;
      this.companyOrder = 0;
      this.capacityOrder = 0;
      this.dateOrder = 1;
    }
    else {
      snapshots.reverse();
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
                        <button className="button is-dark is-small is-pulled-right" onClick={e => this.onSortSerial(e, 'serialNumberInserv')}>
                <i className="fas fa-sort"></i>
              </button>
            </th>
            <th>Company
                        <button className="button is-dark is-small is-pulled-right" onClick={e => this.onSortCompany(e, 'companyName')}>
                <i className="fas fa-sort"></i>
              </button>
            </th>
            <th>Data Used
                        <button className="button is-dark is-small is-pulled-right" onClick={e => this.onSortCapacity(e, 'freeTiB', 'sizeTiB')}>
                <i className="fas fa-sort"></i>
              </button>
            </th>
            <th >Last Updated
                        <button className="button is-dark is-small is-pulled-right" onClick={e => this.onSortDate(e, 'date')}>
                <i className="fas fa-sort"></i>
              </button>
            </th>
          </tr>
        </thead>
        <tbody>
          {
            newData.map(function (system, index) {
              return (
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
                  <td>{system.date}</td>
                </tr>
              );
            })
          }
        </tbody>
      </table>
    )
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
}
