import * as React from "react";
import {ChangeEvent} from "react";

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
                <option key={date}>
                    {date}
                </option>
            );
            items.push(item);
        }

        return (
            <div className="select">
                <select onChange={this.setSnapshotDate}>
                    <option>Latest</option>
                    {items}
                </select>
            </div>
        );
    }
}