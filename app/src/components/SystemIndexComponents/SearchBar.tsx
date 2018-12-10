import * as React from "react";
import '../../sass/custom-bulma.scss';
import {Link} from "react-router-dom";
import {SystemIndexTable} from "./SystemIndexTable";

export class SearchBar extends React.Component{
    render(){
        return(
            <div className="SearchBar is-pulled-right">
                <form action="Search">
                    <input type="text" placeholder="Search.."/>
                    <button type="submit"><i className="fa fa-search"></i></button>
                </form>
            </div>
        )
    }
}