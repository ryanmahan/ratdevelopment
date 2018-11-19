import * as React from "react";
import '../../sass/custom-bulma.scss';
import {Link} from "react-router-dom";


//series of getters to help iteration
function getSerialNumber(currentRow: any){
    return currentRow.serialNumberInserv;
}

function getCompany(currentRow: any){
    return currentRow.system.companyName;
}

function getCapacity(currentRow: any){
    return Math.trunc(100 * (currentRow.capacity.total.freeTiB / currentRow.capacity.total.sizeTiB)) + "%";
}

function getDate(currentRow: any){
    return currentRow.date;
}

//constructs a row, key is to assist React efficiency, calls each of the getters for cells and returns
function createSystemRow(currentRow: any) {
    return <tr key={getSerialNumber(currentRow)}>
            <td>
                <Link to={"/view/" + getSerialNumber(currentRow)} >
                {getSerialNumber(currentRow)}
                </Link>
            </td>
            <td>{getCompany(currentRow)}</td>
            <td>{getCapacity(currentRow)}</td>
            <td>{getDate(currentRow)}</td>
    </tr>
}

//iterates over the array and calls return row for each json. Pushes each to an array and returns it
function populateSystemTable(array: any[]){
    let rows: any[] = [];
    for(let i = 0; i < array.length; i++){
        rows.push(createSystemRow(array[i]));
    }
    return rows;
}




export interface SystemIndexTableProps{

}

interface SystemIndexTableState {
    snapshots: any[]
}

//exports the actual table component which calls our fillArray method
export class SystemIndexTable extends React.Component<SystemIndexTableProps, SystemIndexTableState> {

    constructor(props: SystemIndexTableProps) {
        super(props);
        this.state = {snapshots: []};
    }

    //This is triggered when this component is mounted
    componentDidMount(){
        this.getSnapshots()
    }

    render() {
        return (
            <table id={"myTable"} className="table is-fullwidth is-bordered is-striped">
                <thead>
                <tr>
                    <th>Serial Number</th>
                    <th>Company</th>
                    <th>Capacity Free</th>
                    <th>Last Updated</th>
                </tr>
                </thead>
                <tbody>
                {
                    populateSystemTable(this.state.snapshots)

                }
                </tbody>
            </table>
        )
    }

    //fetch the latest snapshots and then update the state of the table.
    getSnapshots() {
        //  Make the API call
        fetch(
            "http://localhost:8081/GetLatestSnapshotsByTenant?tenant=1200944110"
        ).then(r => {
            //  When that returns convert it to json
            return r.json();
        }).then( j => {
            //  Finally set the state of the table to the list of snapshots returned
            this.setState({
                snapshots: j
            })
        });
    }
}

