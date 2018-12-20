import * as React from "react";
import '../sass/custom-bulma.scss';
import {Divider} from "../components/layout/Divider";
import {PageTitle} from "../components/layout/PageTitle";
import {DateDropdown} from "./SystemView/DateDropdown";
import {match} from "react-router";
import {API_URL} from "../misc/state/constants";
import * as moment from 'moment';

export interface SystemViewProps {
    match: match,
    date: string
}


export interface SystemViewState {
    snapshot: any,
    selectedDate: string,
    validDates: string[]
}

export class SystemView extends React.Component<SystemViewProps, SystemViewState> {

    static defaultProps = {
        date: "latest"
    };

    constructor(props: SystemViewProps){
        super(props);
        this.state = {
            snapshot: {},
            selectedDate: "",
            validDates: [""]
        };
        console.log((this.props.match.params as any).serialNumber);
        this.reload = this.reload.bind(this);
        this.downloadJSON = this.downloadJSON.bind(this);
    }

    componentDidMount(){
        fetch(API_URL + "/api" +
              "/tenants/" +
              "hpe" +
              "/systems/" +
               (this.props.match.params as any).serialNumber +
              "/timestamps"
        ).then(r => {
                return r.json();
            }
        ).then( j =>{
            let date: string;
            date = j[0];
            for (let str in j){
                if (str === this.props.date) {
                    date = this.props.date;
                }
            }

            this.setState({
                selectedDate: date,
                validDates: j
            });
            return fetch(API_URL + "/api" +
                         "/tenants/" +
                         "hpe" +
                         "/systems/" +
                          (this.props.match.params as any).serialNumber +
                         "/snapshots/" +
                         date);
        }).then( r => {
            return r.json();
        }).then( j => {
            this.setState({
                snapshot: j
            })
        }).catch(reason => {
            console.log(reason);
        })
    }

    reload(date: string){
        fetch(API_URL + "/api" +
              "/tenants/" +
              "hpe" +
              "/systems/" +
               (this.props.match.params as any).serialNumber +
              "/snapshots/" +
              date
        ).then( r =>{
                return r.json();
        }).then( j => {
            this.setState({
                snapshot: j,
                selectedDate: date
            });
        })


    }

    render() {
        let serialNumber: string = this.state.snapshot.serialNumberInserv;
        let snapshot: any = this.state.snapshot;
        let date: string = this.state.snapshot.date;
        return (
            <div className="container">
                <PageTitle title={"Serial Number: " + serialNumber}
                           extras={[<DateDropdown reload={this.reload} dates={this.state.validDates} activeDate={this.state.selectedDate}/>]}/>
                <Divider/>
                <div className="level">
                    <div className="level-left title is-5">
                        {moment(date).utc().format('MMMM Do YYYY, h:mm A')}
                    </div>
                    <div className="level-right">
                        <h1></h1>
                        <a className="button level-item is-large" onClick={this.downloadJSON}>
                            Download JSON File &nbsp; <i className="icon fas fa-file-download"/>
                        </a>
                    </div>
                </div>
                <pre className="highlight">
                    <code className="language-json">
                    {JSON.stringify(snapshot, null, 4)}
                    </code>
                </pre>
            </div>
        );
    }

    downloadJSON(){
        let selectedDate = this.state.selectedDate;
        let serialNumber = this.state.snapshot.serialNumberInserv;
        window.location.href = API_URL + "/api" +
                               "/tenants/" +
                               "hpe" +
                               "/systems/" +
                               serialNumber +
                               "/snapshots/" +
                               selectedDate +
                               "/download";
    }
}
