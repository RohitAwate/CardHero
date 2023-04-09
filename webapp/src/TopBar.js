import {Component} from "react";

import "./TopBar.css";


class TopBar extends Component {
    render() {
        return <div id="top-bar" className="non-selectable">
            <p>card</p>
            <p style={{fontWeight: "bold"}}>hero</p>
        </div>;
    }
}

export default TopBar;