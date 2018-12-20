import * as React from "react";
import '../sass/custom-bulma.scss';
import {Divider} from "../components/layout/Divider";
import {PageTitle} from "../components/layout/PageTitle";
import {DateDropdown} from "./SystemView/DateDropdown";
import {match} from "react-router";
import {AppAuthState} from "../misc/state/constants";
let fileDownload = require("js-file-download");
import {API_URL} from "../misc/state/constants";
import * as moment from 'moment';

export interface SystemViewProps {
    match: match,
    date: string,
    authState: AppAuthState
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
              "/timestamps",
            {headers:{Authorization: "BEARER "+this.props.authState.access_token}}
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
                         date,
                {headers:{Authorization: "BEARER "+this.props.authState.access_token}
                });
        }).then( r => {
            return r.json();
        }).then( j => {
            this.setState({
                snapshot: j
            });
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
              date,
            {headers:{Authorization: "BEARER "+this.props.authState.access_token}
            }
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
                    <div className="level-left title is-5" style={{margin: "0"}}>
                        {moment(date).utc().format('MMMM Do YYYY, h:mm A')}
                    </div>
                    <div className="level-right">
                        <a className="button level-item is-large" onClick={this.downloadJSON}>
                            Download JSON File &nbsp; <i className="icon fas fa-file-download"/>
                        </a>
                    </div>
                </div>
                {snapshot.capacity && snapshot.capacity.total && (snapshot.capacity.total.freeTiB / snapshot.capacity.total.sizeTiB <= 0.3) &&
                    <div style={{backgroundColor: "#ffb3b3", color: "#000", padding: "1rem", margin: "0 0 1rem 0"}}>
                        Warning: Free capacity below 30%
                        <figure className="image is-24x24 is-pulled-left" style={{marginRight: "1rem"}}>
                            <img src="https://img.icons8.com/color/50/000000/high-priority.png" alt="Warning: Free capacity below 30%" title="Warning: Free capacity below 30%"></img>
                        </figure>
                    </div>
                }
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

        let xhr: XMLHttpRequest = new XMLHttpRequest();
        xhr.open("GET",
            API_URL +
            "/api" +
            "/tenants/" +
            "hpe" +
            "/systems/" +
            serialNumber +
            "/snapshots/" +
            selectedDate +
            "/download");
        xhr.setRequestHeader('Authorization', "BEARER "+this.props.authState.access_token);
        xhr.onreadystatechange = function()  {
            if (this.readyState == this.DONE) {
                if (this.status === 200) {
                    let json: any = JSON.parse(this.response);
                    fileDownload(this.response, json.serialNumberInserv+"-"+json.date+".json");
                } else {
                    console.error('XHR failed', this);
                }
            }
        };
        xhr.send();
    }
}
