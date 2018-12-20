import * as React from "react";
import * as moment from 'moment';

export interface DateDropdownState {
    activeDate: string,
    dates: string[],
    reload: (date: string) => void
}

export class DateDropdown extends React.Component<DateDropdownState> {

    constructor(props: any){
        super(props);
        this.setSnapshotDate = this.setSnapshotDate.bind(this);
    }

    setSnapshotDate(event: React.ChangeEvent<HTMLSelectElement>){
        let date: string = event.target.value;
        if (date === "Latest") {
            date = this.props.dates[0];
        }
        this.props.reload(date);
    }


    render() {

        let items: any[] = [];
        for (let date of this.props.dates) {
            let item = (
                <option key={date} value={date}>
                    {moment(date).utc().format('MMMM Do YYYY, h:mm A')}
                </option>
            );
            items.push(item);
        }

        return (
            <div className="select level">
                <label style={{minWidth: "4em"}} htmlFor="snapshot-date">Date:</label>
                <select name="snapshot-date" onChange={this.setSnapshotDate}>
                    {items}
                </select>
            </div>
        );
    }
}